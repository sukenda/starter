import { apiRequest } from '@/shared/api/client'
import { useAuthStore } from '@/features/auth/sessionStore'
export interface User { PublicID:string; Email:string; Name:string; Status:string; Roles:string[] }
export interface Role { Code:string; Name:string; Description:string; Permissions:string[]; IsSystem:boolean }
export interface Permission { Code:string; Name:string; Description:string }
export interface Menu { PublicID:string; Code:string; Name:string; Route:{String:string;Valid:boolean}; Icon:{String:string;Valid:boolean}; PermissionCode:{String:string;Valid:boolean}; SortOrder:number; Active:boolean }
function request<T>(path:string,init?:RequestInit){const auth=useAuthStore();return apiRequest<T>(path,{...init,headers:{Authorization:`Bearer ${auth.accessToken}`,...init?.headers}})}
export const adminApi={users:()=>request<User[]>('/api/v1/admin/users'),roles:()=>request<Role[]>('/api/v1/admin/roles'),permissions:()=>request<Permission[]>('/api/v1/admin/permissions'),menus:()=>request<Menu[]>('/api/v1/admin/menus')}
