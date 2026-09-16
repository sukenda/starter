import { defineStore } from 'pinia'
import { apiRequest } from '@/shared/api/client'

export interface CurrentUser { id:string; email:string; name:string; permissions:string[] }
interface Tokens { access_token:string; refresh_token:string; token_type:string; access_expires_at:string; refresh_expires_at:string }
interface LoginResponse { user:Omit<CurrentUser,'permissions'>; tokens:Tokens }

export const useAuthStore=defineStore('auth',{state:()=>({user:null as CurrentUser|null,accessToken:null as string|null,refreshToken:null as string|null,initialized:false}),getters:{isAuthenticated:s=>Boolean(s.user&&s.accessToken),hasPermission:s=>(code:string)=>Boolean(s.user?.permissions.includes(code))},actions:{
 async login(email:string,password:string){const r=await apiRequest<LoginResponse>('/api/v1/auth/login',{method:'POST',body:JSON.stringify({email,password})});this.accessToken=r.tokens.access_token;this.refreshToken=r.tokens.refresh_token;await this.loadMe()},
 async loadMe(){if(!this.accessToken){this.initialized=true;return}try{this.user=await apiRequest<CurrentUser>('/api/v1/auth/me',{headers:{Authorization:`Bearer ${this.accessToken}`}})}catch{this.clear()}finally{this.initialized=true}},
 async logout(){if(this.accessToken){try{await apiRequest('/api/v1/auth/logout',{method:'POST',headers:{Authorization:`Bearer ${this.accessToken}`}})}catch{}}this.clear()},
 clear(){this.user=null;this.accessToken=null;this.refreshToken=null;this.initialized=true},
}})
