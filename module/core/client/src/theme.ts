const keys: string[] = []

export function themeInit() {
  const x = document.querySelectorAll(".color-var");
  if (x.length > 0) {
    for (const el of Array.from(x)) {
      const i = (el as HTMLInputElement)
      const v = i.dataset["var"] as string;
      const m = i.dataset["mode"] as string;
      keys.push(v);
      if (!v || v.length === 0) {
        continue;
      }
      i.oninput = function () {
        set(m, v, i.value);
      }
    }
  }
}

function set(mode: string, key: string, v: string) {
  const mockup = document.querySelector("#mockup-" + mode);
  if (!mockup) {
    console.error("can't find mockup for mode [" + mode + "]")
    return;
  }
  switch (key) {
    case "color-foreground":
      setFG(mockup, ".mock-main", v);
      break;
    case "color-background":
      setBG(mockup, ".mock-main", v);
      break;
    case "color-foreground-muted":
      setFG(mockup, ".mock-main .mock-muted", v);
      break;
    case "color-background-muted":
      setBG(mockup, ".mock-main .mock-muted", v);
      break;
    case "color-link-foreground":
      setFG(mockup, ".mock-main .mock-link", v);
      break;
    case "color-link-visited-foreground":
      setFG(mockup, ".mock-main .mock-link-visited", v);
      break;
    case "color-nav-foreground":
      setFG(mockup, ".mock-nav", v);
      setFG(mockup, ".mock-nav .mock-link", v);
      break;
    case "color-nav-background":
      setBG(mockup, ".mock-nav", v);
      break;
    case "color-menu-foreground":
      setFG(mockup, ".mock-menu", v);
      setFG(mockup, ".mock-menu .mock-link", v);
      break;
    case "color-menu-background":
      setBG(mockup, ".mock-menu", v);
      break;
    case "color-menu-selected-foreground":
      setFG(mockup, ".mock-menu .mock-link-selected", v);
      break;
    case "color-menu-selected-background":
      setBG(mockup, ".mock-menu .mock-link-selected", v);
      break;
    default:
      console.error("invalid key [" + key + "]")
  }
}

function call(mockup: Element, sel: string, f: (el: HTMLElement) => void) {
  const q = mockup.querySelectorAll(sel);
  if (q.length == 0) {
    throw "empty query selector [" + sel + "]"
  }
  q.forEach(x => f(x as HTMLElement));
}

function setBG(mockup: Element, sel: string, v: string) {
  call(mockup, sel, el => el.style.backgroundColor = v);
}

function setFG(mockup: Element, sel: string, v: string) {
  call(mockup, sel, el => el.style.color = v);
}
