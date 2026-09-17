<script setup lang="ts">
import { reactive, watch } from 'vue'
import DsButton from '@/shared/ui/DsButton.vue'
import DsDialog from '@/shared/ui/DsDialog.vue'
import DsInput from '@/shared/ui/DsInput.vue'
import DsSelect from '@/shared/ui/DsSelect.vue'
import type { Role, User, UserInput } from '../api'

const p = defineProps<{ open: boolean; user?: User | null; roles: Role[]; busy?: boolean }>()
const emit = defineEmits<{ close: []; save: [UserInput] }>()
const f = reactive({ email: '', name: '', password: '', status: 'active', roles: [] as string[] })
watch(
  () => [p.open, p.user] as const,
  () => {
    f.email = p.user?.Email ?? ''
    f.name = p.user?.Name ?? ''
    f.password = ''
    f.status = p.user?.Status ?? 'active'
    f.roles = [...(p.user?.Roles ?? [])]
  },
  { immediate: true },
)
function toggle(code: string) {
  f.roles = f.roles.includes(code) ? f.roles.filter((value) => value !== code) : [...f.roles, code]
}
function save() {
  const input: UserInput = { email: f.email, name: f.name, status: f.status, roles: [...f.roles] }
  if (!p.user) input.password = f.password
  emit('save', input)
}
</script>
<template>
  <DsDialog :open="open" :title="user ? 'Edit user' : 'Add user'" description="Identity and role assignments are enforced by the backend." @close="emit('close')">
    <form class="form" @submit.prevent="save">
      <DsInput v-model="f.name" label="Name" />
      <DsInput v-model="f.email" label="Email" type="email" />
      <DsInput v-if="!user" v-model="f.password" label="Temporary password" type="password" />
      <DsSelect v-model="f.status" label="Status" :options="[{ label: 'Active', value: 'active' }, { label: 'Disabled', value: 'disabled' }]" />
      <fieldset>
        <legend>Roles</legend>
        <label v-for="r in roles" :key="r.Code"><input type="checkbox" :checked="f.roles.includes(r.Code)" @change="toggle(r.Code)">{{ r.Name }}</label>
      </fieldset>
    </form>
    <template #footer>
      <DsButton variant="secondary" @click="emit('close')">Cancel</DsButton>
      <DsButton :disabled="busy || !f.name || !f.email || (!user && f.password.length < 12)" @click="save">{{ busy ? 'Saving…' : 'Save user' }}</DsButton>
    </template>
  </DsDialog>
</template>
<style scoped>.form{display:grid;gap:16px}fieldset{display:grid;gap:10px;margin:0;padding:14px;border:1px solid var(--ds-border);border-radius:var(--ds-radius-md)}legend{padding:0 5px;font-size:13px;font-weight:600}fieldset label{display:flex;gap:9px;align-items:center;font-size:14px}</style>
