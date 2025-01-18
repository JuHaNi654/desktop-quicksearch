import './style.css';
import './app.css';
import PlaceholderIcon from './assets/icon.svg'
import {Search, RunCommand} from '../wailsjs/go/main/DesktopEntries';
import {Exit} from '../wailsjs/go/main/App';

function start(exec: string) {
  try {
    RunCommand(exec)
  } catch (err) {
    console.error(err)
  } 
}

function quit() {
  try {
    Exit() 
  } catch (err) {
    console.error(err)
  } 
}

function search(e: Event, results: HTMLUListElement) { 
  const target = e.target as HTMLInputElement 
  let pattern = target.value;
  if (pattern === "" || !pattern) return;

  results.innerHTML = ""
  try {
    Search(pattern)
      .then(async (result) => {
        const data = JSON.parse(result)
        if (data.length === 0) return
        for (let i = 0; i < data.length; i++) {
          const item = document.createElement("li")
          const button = document.createElement("button")
          button.addEventListener("click", () => start(data[i].desktop))
         
          const icon = new Image(40, 40)
          icon.alt = `${data[i].name} icon`
          if (data[i].icon && data[i].icon !== "") {  
            const img = await fetch(data[i].icon)
            if (img.status === 200) {
              const blob = await img.blob()
              icon.src = URL.createObjectURL(blob)
            } else {
              icon.src = PlaceholderIcon
            }
          } else {
            icon.src = PlaceholderIcon
          }

          button.appendChild(icon)
          const label = document.createElement("span")
          label.innerText = data[i].name
          button.appendChild(label)
          item.appendChild(button)
          results.appendChild(item)
        } 
      })
  } catch (err) {
    console.error(err)
  } 
}

let keyMap: { [key: number]: boolean } = {}
function handleKeyInput(e: KeyboardEvent) {
  if (e.type === "keydown") {
    keyMap[e.keyCode] = true
    if (keyMap[27]) {
      quit() 
    } else if (keyMap[17] && keyMap[67]) {
      quit()
    }
  }

  if (e.type === "keyup") {
    delete keyMap[e.keyCode]
  }
}

document.addEventListener("DOMContentLoaded", () => {
  const app = document.getElementById("app") as HTMLDivElement

  const container = document.createElement("div")
  container.setAttribute("class", "input-box")
  container.setAttribute("id", "input")
  app.appendChild(container)

  const results = document.createElement("ul")
  results.setAttribute("class", "results")

  const searchInput = document.createElement("input")
  searchInput.setAttribute("class", "input")
  searchInput.setAttribute("id", "search")
  searchInput.setAttribute("type", "text")
  searchInput.setAttribute("autocomplete", "off")
  searchInput.addEventListener("input", (e) => search(e, results))

  container.appendChild(searchInput)
  container.appendChild(results)
  searchInput.focus();


  window.addEventListener("keydown", handleKeyInput)
  window.addEventListener("keyup", handleKeyInput)
  window.addEventListener("blur", () => keyMap = {})
})


//let resultElement = document.getElementById("result");

