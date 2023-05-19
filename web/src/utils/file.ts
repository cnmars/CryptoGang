export const fileUtils ={
    readFile: async (file:any)=> {
        const reader = new FileReader()
        const promise = new Promise((resolve, reject) => {
          reader.onload = function () {
            resolve(reader.result)
          }
          reader.onerror = function (e) {
            reader.abort()
            reject(e)
          }
        })
        reader.readAsText(file.raw, 'UTF-8') // 将文件读取为文本
        return promise
      }
}