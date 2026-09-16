<script setup lang="ts">
import { useAuthStore } from '@/features/auth/sessionStore'
const auth=useAuthStore()
const items=[['Dashboard','/'],['Users','/users','users.read'],['Roles','/roles','roles.read'],['Menus','/menus','menus.read']] as const
</script>
<template><div class="shell"><aside><strong>Starter Admin</strong><nav><template v-for="i in items" :key="i[1]"><RouterLink v-if="!i[2]||auth.hasPermission(i[2])" :to="i[1]">{{i[0]}}</RouterLink></template></nav></aside><main><header><span>{{auth.user?.name}}</span><button @click="auth.logout()">Logout</button></header><RouterView/></main></div></template>
<style scoped>.shell{min-height:100vh;display:grid;grid-template-columns:240px 1fr;background:#f7f8fa;color:#171717;font-family:Inter,system-ui,sans-serif}aside{padding:28px;background:#fff;border-right:1px solid #e5e7eb}nav{display:grid;gap:8px;margin-top:28px}a{padding:10px 12px;color:#525252;text-decoration:none;border-radius:8px}.router-link-active{background:#f1f5f9;color:#111827}main{min-width:0}header{height:64px;background:#fff;border-bottom:1px solid #e5e7eb;display:flex;align-items:center;justify-content:flex-end;gap:16px;padding:0 28px}button{padding:8px 12px;border:1px solid #d1d5db;background:#fff;border-radius:8px;cursor:pointer}@media(max-width:720px){.shell{grid-template-columns:1fr}aside{display:none}}</style>
