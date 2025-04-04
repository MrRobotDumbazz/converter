<template>
    <div class="container mx-auto p-6 max-w-2xl">
      <div class="bg-white shadow-md rounded-lg p-8">
        <h1 class="text-3xl font-bold text-center mb-6 text-gray-800">
          Universal Converter
        </h1>
  
        <div class="grid gap-4">
          <div class="grid md:grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                Category
              </label>
              <select 
                v-model="selectedCategory" 
                @change="resetUnits"
                class="w-full p-3 border border-gray-300 rounded-md"
              >
                <option 
                  v-for="(category, key) in categories" 
                  :key="key" 
                  :value="key"
                >
                  {{ category }}
                </option>
              </select>
            </div>
  
            <div class="grid grid-cols-2 gap-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">
                  From
                </label>
                <select 
                  v-model="fromUnit" 
                  class="w-full p-3 border border-gray-300 rounded-md"
                >
                  <option 
                    v-for="unit in currentUnits" 
                    :key="unit" 
                    :value="unit"
                  >
                    {{ unit }}
                  </option>
                </select>
              </div>
  
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">
                  To
                </label>
                <select 
                  v-model="toUnit" 
                  class="w-full p-3 border border-gray-300 rounded-md"
                >
                  <option 
                    v-for="unit in currentUnits" 
                    :key="unit" 
                    :value="unit"
                  >
                    {{ unit }}
                  </option>
                </select>
              </div>
            </div>
          </div>
  
          <div class="grid md:grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                Value
              </label>
              <input 
                v-model.number="inputValue" 
                type="number" 
                class="w-full p-3 border border-gray-300 rounded-md"
                placeholder="Enter value"
              >
            </div>
  
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                Result
              </label>
              <div class="w-full p-3 border border-gray-300 rounded-md bg-gray-50">
                <span class="float-right font-semibold">
                  {{ result !== null ? result.toFixed(6) : '—' }} 
                  {{ result !== null ? toUnit : '' }}
                </span>
              </div>
            </div>
          </div>
  
          <div class="flex justify-center mt-4">
            <button 
              @click="convert" 
              class="bg-blue-500 text-white px-6 py-3 rounded-md hover:bg-blue-600 transition"
            >
              Convert
            </button>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script lang="ts">
  import { defineComponent } from 'vue'
  import * as App from '../../wailsjs/go/main/App'
  
  export default defineComponent({
    name: 'Converter',
    data() {
      return {
        categories: {
          length: ['nm', 'mcm', 'mm', 'cm', 'dm', 'm', 'km'],
          data: ['bit', 'byte', 'KiB', 'MiB', 'GiB', 'TiB', 'PiB'],
          speed: ['bit/s', 'byte/s', 'Kbit/s', 'Mbit/s', 'Gbit/s'],
          time: ['ns', 'mcs', 'ms', 'sec', 'min', 'hour', 'day', 'week', 'month', 'year'],
          weight: ['mcg', 'mg', 'g', 'kg', 'ton']
        },
        selectedCategory: 'length',
        fromUnit: 'm',
        toUnit: 'km',
        inputValue: 1,
        result: null as number | null
      }
    },
    computed: {
      currentUnits(): string[] {
        return this.categories[this.selectedCategory as keyof typeof this.categories]
      }
    },
    methods: {
      async convert() {
        try {
          this.result = await App.Convert(
            this.inputValue, 
            this.fromUnit, 
            this.toUnit
          )
        } catch (error) {
          console.error('Conversion error:', error)
          this.result = null
        }
      },
      resetUnits() {
        const units = this.currentUnits
        this.fromUnit = units[0]
        this.toUnit = units[1]
      }
    },
    mounted() {
      this.resetUnits()
    }
  })
  </script>