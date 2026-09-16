<script setup lang="ts">
import { useQuery } from '@tanstack/vue-query'

interface HealthResponse {
  status: string
}

const health = useQuery({
  queryKey: ['health'],
  queryFn: async ({ signal }): Promise<HealthResponse> => {
    const response = await fetch('/health', { signal })
    if (!response.ok) throw new Error('Health check failed')
    return response.json() as Promise<HealthResponse>
  },
})
</script>

<template>
  <main>
    <h1>Fullstack Starter</h1>
    <p v-if="health.isPending.value">Checking backend…</p>
    <p v-else-if="health.isError.value">Backend unavailable</p>
    <p v-else>Backend: {{ health.data.value?.status }}</p>
  </main>
</template>
