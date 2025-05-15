<template>
    <button @click='onIconClick' :class="['inline-flex items-center justify-center transition']">
        <span v-html="iconContent" :class="iconClass" />
    </button>
</template>

<script lang="ts">
import { defineComponent, ref, watch } from 'vue'

export default defineComponent({
    name: 'IconButton',
    props: {
        name: {
            type: String,
            required: true
        },
        iconClass: {
            type: String,
            default: 'h-5 w-5'
        }
    },

    emits: ['click'],
    data() {
        return {
            iconContent: ''
        }
    },
    methods: {
        onIconClick() {
            this.$emit('click')
        },
        async loadIcon() {
            const icons = import.meta.glob('../../assets/icons/*.svg', { as: 'raw' })
            const importIcon = icons[`../../assets/icons/${this.name}.svg`]
            if (!importIcon) {
                console.warn(`Icon ${this.name} not found!`)
                this.iconContent = ''
                return
            }
            this.iconContent = await importIcon()
        }
    },
    watch: {
        name: {
            immediate: true,
            handler() {
                this.loadIcon()
            }
        }
    }
})
</script>

<style scoped>
/* Tuỳ chỉnh kích thước icon */
.h-5 {
    width: 1.25rem;
    height: 1.25rem;
    display: inline-block;
    vertical-align: middle;
}
</style>