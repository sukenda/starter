<script setup lang="ts">
import { ref } from 'vue'
import { useMutation, useQuery, useQueryClient } from '@tanstack/vue-query'
import { adminApi, type Menu, type MenuInput } from '../api'
import MenuDialog from '../components/MenuDialog.vue'
import DsBadge from '@/shared/ui/DsBadge.vue'
import DsButton from '@/shared/ui/DsButton.vue'
import DsCard from '@/shared/ui/DsCard.vue'
import DsPageHeader from '@/shared/ui/DsPageHeader.vue'
import DsState from '@/shared/ui/DsState.vue'

const qc = useQueryClient()
const selected = ref<Menu | null>(null)
const open = ref(false)
const q = useQuery({ queryKey: ['admin', 'menus'], queryFn: adminApi.menus })
const permissions = useQuery({ queryKey: ['admin', 'permissions'], queryFn: adminApi.permissions })
const save = useMutation({
  mutationFn: (value: MenuInput) => selected.value ? adminApi.updateMenu(selected.value.PublicID, value) : adminApi.createMenu(value),
  onSuccess: async () => {
    open.value = false
    selected.value = null
    await qc.invalidateQueries({ queryKey: ['admin', 'menus'] })
  },
})
function edit(menu: Menu) {
  selected.value = menu
  open.value = true
}
</script>
<template>
  <section>
    <DsPageHeader eyebrow="NAVIGATION" title="Menu Management" description="Configure application navigation metadata and permission visibility.">
      <template #actions><DsButton @click="selected = null; open = true">Add menu</DsButton></template>
    </DsPageHeader>
    <DsCard class="panel">
      <DsState v-if="q.isPending.value" kind="loading" />
      <DsState v-else-if="q.isError.value" kind="error" title="Unable to load menus" />
      <DsState v-else-if="!q.data.value?.length" kind="empty" title="No menus configured" />
      <table v-else>
        <thead><tr><th>Menu</th><th>Code</th><th>Route</th><th>Permission</th><th>Status</th><th></th></tr></thead>
        <tbody><tr v-for="m in q.data.value" :key="m.PublicID"><td><strong>{{ m.Name }}</strong></td><td><code>{{ m.Code }}</code></td><td>{{ m.Route || '—' }}</td><td>{{ m.PermissionCode || '—' }}</td><td><DsBadge :tone="m.Active ? 'success' : 'neutral'">{{ m.Active ? 'Active' : 'Inactive' }}</DsBadge></td><td><DsButton variant="secondary" @click="edit(m)">Edit</DsButton></td></tr></tbody>
      </table>
    </DsCard>
    <MenuDialog :open="open" :menu="selected" :permissions="permissions.data.value ?? []" :busy="save.isPending.value" @close="open = false" @save="save.mutate" />
  </section>
</template>
<style scoped>section{padding:32px}.panel{padding:0;overflow:auto}table{width:100%;border-collapse:collapse}th,td{text-align:left;padding:15px 18px;border-bottom:1px solid var(--ds-border);font-size:14px}th{color:var(--ds-text-muted);font-size:12px;text-transform:uppercase;letter-spacing:.05em}tbody tr:last-child td{border-bottom:0}code{color:var(--ds-text-secondary)}td:last-child{text-align:right}@media(max-width:640px){section{padding:20px}}</style>
