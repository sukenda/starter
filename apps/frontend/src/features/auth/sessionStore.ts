import { defineStore } from 'pinia'
import { apiRequest } from '@/shared/api/client'

export interface CurrentUser { id:string; email:string; name:string; permissions:string[] }
interface Tokens { access_token:string; token_type:string; access_expires_at:string; refresh_expires_at:string }
interface LoginResponse { user:Omit<CurrentUser,'permissions'>; tokens:Tokens }

function csrfToken(){const item=document.cookie.split('; ').find(v=>v.startsWith('starter_csrf='));return item?decodeURIComponent(item.slice('starter_csrf='.length)):''}

export const useAuthStore=defineStore('auth',{state:()=>({user:null as CurrentUser|null,accessToken:null as string|null,initialized:false,refreshing:null as Promise<boolean>|null}),getters:{isAuthenticated:s=>Boolean(s.user&&s.accessToken),hasPermission:s=>(code:string)=>Boolean(s.user?.permissions.includes(code))},actions:{
 async login(email:string,password:string){const r=await apiRequest<LoginResponse>('/api/v1/auth/login',{method:'POST',credentials:'include',body:JSON.stringify({email,password,client:'web'})});this.accessToken=r.tokens.access_token;await this.loadMe()},
 async refresh(){if(this.refreshing)return this.refreshing;this.refreshing=(async()=>{try{const csrf=csrfToken();const r=await apiRequest<Tokens>('/api/v1/auth/refresh',{method:'POST',credentials:'include',headers:{'X-CSRF-Token':csrf},body:'{}'});this.accessToken=r.access_token;return true}catch{this.clear();return false}finally{this.refreshing=null}})();return this.refreshing},
 async loadMe(){if(!this.accessToken){if(!(await this.refresh())){this.initialized=true;return}}try{this.user=await apiRequest<CurrentUser>('/api/v1/auth/me',{headers:{Authorization:`Bearer ${this.accessToken}`}})}catch{if(await this.refresh()){try{this.user=await apiRequest<CurrentUser>('/api/v1/auth/me',{headers:{Authorization:`Bearer ${this.accessToken}`}})}catch{this.clear()}}else this.clear()}finally{this.initialized=true}},
 async logout(){if(this.accessToken){try{await apiRequest('/api/v1/auth/logout',{method:'POST',credentials:'include',headers:{Authorization:`Bearer ${this.accessToken}`}})}catch{}}this.clear()},
 clear(){this.user=null;this.accessToken=null;this.initialized=true},
}})
