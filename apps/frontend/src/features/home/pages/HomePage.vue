<script setup lang="ts">
const stack = ['Vue 3', 'TypeScript', 'Vite', 'TanStack Query', 'Pinia']
</script>

<template>
  <main>
    <h1>Fullstack Starter</h1>
    <p>Frontend is ready for feature development.</p>
    <ul>
      <li v-for="item in stack" :key="item">{{ item }}</li>
    </ul>
  </main>
</template>
