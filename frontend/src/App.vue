<template>
  <div class="h-screen flex flex-col bg-gray-500">
    <!-- Top Bar -->
    <div class="flex justify-center">
      <TopBar />
    </div>

    <!-- Main Layout -->
    <div class="flex-1 justify-center">
      <!-- Main Content -->
      <main class="flex px-4 w-full justify-center">
        <router-view />
      </main>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import TopBar from './components/TopBar.vue'

export default defineComponent({
  name: 'App',
  components: { TopBar },
  data() {
    return {
      isSidebarOpen: false
    }
  }
})
</script>

<style scoped>
aside {
  width: 300px;
  top: 64px;
  bottom: 0;
}

@media (min-width: 768px) {
  aside {
    top: 0;
  }
}
</style>