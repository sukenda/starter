import { apiRequest } from '@/shared/api/client'
import { useAuthStore } from '@/features/auth/sessionStore'

export interface User { PublicID:string; Email:string; Name:string; Status:string; Roles:string[] }
export interface Role { Code:string; Name:string; Description:string; Permissions:string[]; IsSystem:boolean }
export interface Permission { Code:string; Name:string; Description:string }
export interface Menu { PublicID:string; ParentID:string|null; Code:string; Name:string; Route:string|null; Icon:string|null; PermissionCode:string|null; SortOrder:number; Active:boolean }
export interface UserInput { email:string;name:string;password?:string;status:string;roles:string[] }
export interface RoleInput { code?:string;name:string;description:string;permissions:string[] }
export interface MenuInput { code?:string;name:string;route:string|null;icon:string|null;permission:string|null;sort_order:number;active:boolean;parent_id:string|null }

type UserDTO={id:string;email:string;name:string;status:string;roles:string[]}
type RoleDTO={code:string;name:string;description:string;permissions:string[];system:boolean}
type PermissionDTO={code:string;name:string;description:string}
type MenuDTO={id:string;parent_id:string|null;code:string;name:string;route:string|null;icon:string|null;permission:string|null;sort_order:number;active:boolean}

function request<T>(path:string,init?:RequestInit){const auth=useAuthStore();return apiRequest<T>(path,{...init,headers:{Authorization:`Bearer ${auth.accessToken}`,...init?.headers}})}
const json=(method:string,body:unknown):RequestInit=>({method,body:JSON.stringify(body)})
const user=(x:UserDTO):User=>({PublicID:x.id,Email:x.email,Name:x.name,Status:x.status,Roles:x.roles})
const role=(x:RoleDTO):Role=>({Code:x.code,Name:x.name,Description:x.description,Permissions:x.permissions,IsSystem:x.system})
const permission=(x:PermissionDTO):Permission=>({Code:x.code,Name:x.name,Description:x.description})
const menu=(x:MenuDTO):Menu=>({PublicID:x.id,ParentID:x.parent_id,Code:x.code,Name:x.name,Route:x.route,Icon:x.icon,PermissionCode:x.permission,SortOrder:x.sort_order,Active:x.active})

export const adminApi={
 users:async()=> (await request<UserDTO[]>('/api/v1/admin/users')).map(user),
 createUser:(v:UserInput)=>request<void>('/api/v1/admin/users',json('POST',v)),
 updateUser:(id:string,v:UserInput)=>request<void>(`/api/v1/admin/users/${id}`,json('PUT',v)),
 roles:async()=> (await request<RoleDTO[]>('/api/v1/admin/roles')).map(role),
 permissions:async()=> (await request<PermissionDTO[]>('/api/v1/admin/permissions')).map(permission),
 createRole:(v:RoleInput)=>request<void>('/api/v1/admin/roles',json('POST',v)),
 updateRole:(code:string,v:RoleInput)=>request<void>(`/api/v1/admin/roles/${code}`,json('PUT',v)),
 menus:async()=> (await request<MenuDTO[]>('/api/v1/admin/menus')).map(menu),
 createMenu:(v:MenuInput)=>request<void>('/api/v1/admin/menus',json('POST',v)),
 updateMenu:(id:string,v:MenuInput)=>request<void>(`/api/v1/admin/menus/${id}`,json('PUT',v))
}
