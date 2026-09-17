<script setup lang="ts">
import { reactive, watch } from 'vue'
import DsButton from '@/shared/ui/DsButton.vue'
import DsDialog from '@/shared/ui/DsDialog.vue'
import DsInput from '@/shared/ui/DsInput.vue'
import type { Permission, Role, RoleInput } from '../api'

const p = defineProps<{ open: boolean; role?: Role | null; permissions: Permission[]; busy?: boolean }>()
const emit = defineEmits<{ close: []; save: [RoleInput] }>()
const f = reactive({ code: '', name: '', description: '', permissions: [] as string[] })
watch(
  () => [p.open, p.role] as const,
  () => {
    f.code = p.role?.Code ?? ''
    f.name = p.role?.Name ?? ''
    f.description = p.role?.Description ?? ''
    f.permissions = [...(p.role?.Permissions ?? [])]
  },
  { immediate: true },
)
function toggle(code: string) {
  f.permissions = f.permissions.includes(code) ? f.permissions.filter((value) => value !== code) : [...f.permissions, code]
}
function save() {
  emit('save', {
    code: p.role ? undefined : f.code,
    name: f.name,
    description: f.description,
    permissions: [...f.permissions],
  })
}
</script>
<template>
  <DsDialog :open="open" :title="role ? 'Edit role' : 'Add role'" description="Assign explicit backend permissions to this role." @close="emit('close')">
    <div class="form">
      <DsInput v-if="!role" v-model="f.code" label="Code" placeholder="operations_manager" />
      <DsInput v-model="f.name" label="Name" />
      <DsInput v-model="f.description" label="Description" />
      <fieldset>
        <legend>Permissions</legend>
        <label v-for="x in permissions" :key="x.Code"><input type="checkbox" :checked="f.permissions.includes(x.Code)" @change="toggle(x.Code)"><span><strong>{{ x.Name }}</strong><small>{{ x.Code }}</small></span></label>
      </fieldset>
    </div>
    <template #footer>
      <DsButton variant="secondary" @click="emit('close')">Cancel</DsButton>
      <DsButton :disabled="busy || !f.name || (!role && !f.code)" @click="save">{{ busy ? 'Saving…' : 'Save role' }}</DsButton>
    </template>
  </DsDialog>
</template>
<style scoped>.form{display:grid;gap:16px}fieldset{display:grid;gap:8px;margin:0;padding:14px;border:1px solid var(--ds-border);border-radius:var(--ds-radius-md);max-height:280px;overflow:auto}legend{font-size:13px;font-weight:600}fieldset label{display:flex;gap:10px;align-items:flex-start;padding:6px}fieldset span{display:grid}small{color:var(--ds-text-muted)}</style>
