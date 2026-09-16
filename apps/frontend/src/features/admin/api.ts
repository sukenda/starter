import { apiRequest } from '@/shared/api/client'
import { useAuthStore } from '@/features/auth/sessionStore'
export interface User { PublicID:string; Email:string; Name:string; Status:string; Roles:string[] }
export interface Role { Code:string; Name:string; Description:string; Permissions:string[]; IsSystem:boolean }
export interface Permission { Code:string; Name:string; Description:string }
export interface Menu { PublicID:string; Code:string; Name:string; Route:{String:string;Valid:boolean}; Icon:{String:string;Valid:boolean}; PermissionCode:{String:string;Valid:boolean}; SortOrder:number; Active:boolean }
export interface UserInput { email:string;name:string;password?:string;status:string;roles:string[] }
export interface RoleInput { code?:string;name:string;description:string;permissions:string[] }
export interface MenuInput { code?:string;name:string;route:string|null;icon:string|null;permission:string|null;sort_order:number;active:boolean;parent_id:number|null }
function request<T>(path:string,init?:RequestInit){const auth=useAuthStore();return apiRequest<T>(path,{...init,headers:{Authorization:`Bearer ${auth.accessToken}`,...init?.headers}})}
const json=(method:string,body:unknown):RequestInit=>({method,body:JSON.stringify(body)})
export const adminApi={users:()=>request<User[]>('/api/v1/admin/users'),createUser:(v:UserInput)=>request<void>('/api/v1/admin/users',json('POST',v)),updateUser:(id:string,v:UserInput)=>request<void>(`/api/v1/admin/users/${id}`,json('PUT',v)),roles:()=>request<Role[]>('/api/v1/admin/roles'),permissions:()=>request<Permission[]>('/api/v1/admin/permissions'),createRole:(v:RoleInput)=>request<void>('/api/v1/admin/roles',json('POST',v)),updateRole:(code:string,v:RoleInput)=>request<void>(`/api/v1/admin/roles/${code}`,json('PUT',v)),menus:()=>request<Menu[]>('/api/v1/admin/menus'),createMenu:(v:MenuInput)=>request<void>('/api/v1/admin/menus',json('POST',v)),updateMenu:(id:string,v:MenuInput)=>request<void>(`/api/v1/admin/menus/${id}`,json('PUT',v))}
