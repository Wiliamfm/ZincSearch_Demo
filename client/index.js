import { createApp, reactive } from "https://unpkg.com/vue@3/dist/vue.esm-browser.js"
document.addEventListener("DOMContentLoaded", () => {
   console.log("Working")
   const app = createApp({
      setup() {
         const counter = reactive({ count: 0 })
         return {
            counter
         }
      }
   })

   app.mount('#app')
})