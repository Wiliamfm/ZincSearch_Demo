import { createApp, ref, reactive } from "https://unpkg.com/vue@3/dist/vue.esm-browser.js";

document.addEventListener("DOMContentLoaded", () => {
   console.log("Working")
   const app = createApp({
      setup() {
         const options = ref([
            { value: "any", text: "Any" },
            { value: "from", text: "From" },
            { value: "to", text: "To" },
            { value: "subject", text: "Subject" },
         ])
         const selected = ref("any")
         const inputSearch = ref("")
         const data = ref([])
         const headers = ref([
            "From",
            "To",
            "Subject",
            //"Content",
            "Original Email"
         ]);
         const pageClass = ref("")

         function submit(e) {
            e.preventDefault()
            data.value = {
               search: inputSearch.value,
               type: selected.value
            }
            makeSearch(data.value)
               .then(response => response.json())
               .then(search => {
                  inputSearch.value = "";
                  selected.value = "any"
                  data.value = parseContent(search["hits"]["hits"]);
               })
               .catch(err => {
                  console.error("Could not make search: ", err)
               })
         }

         return {
            options,
            selected,
            inputSearch,
            headers,
            data,
            submit
         }
      }
   })

   app.mount('#body')
})

function makeSearch(data) {
   return fetch("http://localhost:3000/zincsearch", {
      method: "post",
      headers: {
         "Content-Type": "application/json",
      },
      body: JSON.stringify(data)
   })
}

function parseContent(emails) {
   const fromRegex = /^From: .*$/m;
   const toRegex = /^To: .*$/m;
   const subjectRegex = /^Subject: .*$/m;
   const headersRegex = /^\w+(-\w+)*:\s+.*$/gm;
   const re = /^\w+: /;
   let parseEmails = [];
   emails.forEach(content => {
      let msg = content["_source"]["Content"];
      let from = fromRegex.exec(msg);
      let to = toRegex.exec(msg);
      let subject = subjectRegex.exec(msg);
      //let headers = headersRegex.exec(msg)
      if (from == null || from.length != 1) {
         from = ["null: null"]
      }
      if (to == null || to.length != 1) {
         to = ["null: null"]
      }
      if (subject == null || subject.length != 1) {
         subject = ["null: null"]
      }
      msg = msg.replaceAll(/\\n{2,}/g, "");
      console.log(msg);
      //msg = msg.replaceAll(headersRegex, "").trim()
      parseEmails.push({
         from: re[Symbol.replace](from[0], ""),
         to: re[Symbol.replace](to[0], ""),
         subject: re[Symbol.replace](subject[0], ""),
         content: msg
      })
   });
   return parseEmails;
}