(()=>{function y(){for(let t of Array.from(document.querySelectorAll(".menu-container .final")))t.scrollIntoView({block:"nearest"})}var g="mode-light",k="mode-dark";function v(){for(let t of Array.from(document.getElementsByClassName("mode-input"))){let n=t;n.onclick=function(){switch(n.value){case"":document.body.classList.remove(g),document.body.classList.remove(k);break;case"light":document.body.classList.add(g),document.body.classList.remove(k);break;case"dark":document.body.classList.remove(g),document.body.classList.add(k);break}}}}function b(){let t=document.getElementById("flash-container");if(t===null)return;let n=t.querySelectorAll(".flash");n.length>0&&setTimeout(()=>{for(let e of n){let o=e;o.style.opacity="0",setTimeout(()=>o.remove(),500)}},3e3)}function L(){for(let t of Array.from(document.getElementsByClassName("link-confirm"))){let n=t;n.onclick=function(){let e=n.dataset.message;return e&&e.length===0&&(e="Are you sure?"),confirm(e)}}}function M(){document.addEventListener("keydown",t=>{t.key==="Escape"&&document.location.hash.startsWith("#modal-")&&(document.location.hash="")})}function f(t,n){let e;n?e=n.querySelectorAll(t):e=document.querySelectorAll(t);let o=[];return e.forEach(r=>{o.push(r)}),o}function h(t,n){let e=f(t,n);switch(e.length){case 0:return;case 1:return e[0];default:console.warn(`found [${e.length}] elements with selector [${t}], wanted zero or one`)}}function a(t,n){let e=h(t,n);return e||console.warn(`no element found for selector [${t}]`),e}function l(t,n,e="block"){return typeof t=="string"&&(t=a(t)),t.style.display=n?e:"none",t}function T(t,n){return`<svg class="icon" style="width: ${n}px; height: ${n}px;"><use xlink:href="#svg-${t}"></use></svg>`}function H(){for(let t of f(".tag-editor")){let n=a("input.result",t),e=a(".tags",t),o=n.value.split(",").map(c=>c.trim()).filter(c=>c!=="");l(n,!1),e.innerHTML="";for(let c of o)e.appendChild(x(c,t));h(".add-item",t)?.remove();let r=document.createElement("div");r.className="add-item",r.innerHTML=T("plus",22),r.onclick=function(){B(e,t)},t.insertBefore(r,a(".clear",t))}}function $(t,n){return t.parentElement!==n.parentElement?null:t===n?0:t.compareDocumentPosition(n)&Node.DOCUMENT_POSITION_FOLLOWING?-1:1}var p;function x(t,n){let e=document.createElement("div");e.className="item",e.draggable=!0,e.ondragstart=function(s){s.dataTransfer.setDragImage(document.createElement("div"),0,0),e.classList.add("dragging"),p=e},e.ondragover=function(s){let m=$(e,p);if(!m)return;let u=m===-1?e:e.nextSibling;p.parentElement.insertBefore(p,u),E(n)},e.ondrop=function(s){s.preventDefault()},e.ondragend=function(s){e.classList.remove("dragging"),s.preventDefault()};let o=document.createElement("div");o.innerText=t,o.className="value",o.onclick=function(){I(e)},e.appendChild(o);let r=document.createElement("input");r.className="editor",e.appendChild(r);let c=document.createElement("div");return c.innerHTML=T("times",13),c.className="close",c.onclick=function(){j(e)},e.appendChild(c),e}function j(t){let n=t.parentElement.parentElement;t.remove(),E(n)}function B(t,n){let e=x("",n);t.appendChild(e),I(e)}function I(t){let n=a(".value",t),e=a(".editor",t);e.value=n.innerText;let o=function(){if(e.value===""){t.remove();return}n.innerText=e.value,l(n,!0),l(e,!1);let r=t.parentElement.parentElement;E(r)};e.onblur=o,e.onkeydown=function(r){if(r.code==="Enter")return r.preventDefault(),o(),!1},l(n,!1),l(e,!0),e.focus()}function E(t){let n=[],e=f(".item .value",t);for(let r of e)n.push(r.innerText);let o=a("input.result",t);o.value=n.join(", ")}var w="--selected";function W(t){let n=t.parentElement.parentElement.querySelector("input");if(!n)throw"no associated input found";n.value="\u2205"}function S(){window.projectforge.setSiblingToNull=W;let t={},n={};for(let e of Array.from(document.querySelectorAll("form.editor"))){let o=e,r=()=>{t={},n={};for(let c of o.elements){let s=c;if(s.name.length>0)if(s.name.endsWith(w))n[s.name]=s;else{(s.type!=="radio"||s.checked)&&(t[s.name]=s.value);let m=()=>{let u=n[s.name+w];u&&(u.checked=t[s.name]!==s.value)};s.onchange=m,s.onkeyup=m}}};o.onreset=r,r()}}var F=[];function N(){let t=document.querySelectorAll(".color-var");if(t.length>0)for(let n of Array.from(t)){let e=n,o=e.dataset.var,r=e.dataset.mode;F.push(o),!(!o||o.length===0)&&(e.oninput=function(){R(r,o,e.value)})}}function R(t,n,e){let o=document.querySelector("#mockup-"+t);if(!o){console.error("can't find mockup for mode ["+t+"]");return}switch(n){case"color-foreground":i(o,".mock-main",e);break;case"color-background":d(o,".mock-main",e);break;case"color-foreground-muted":i(o,".mock-main .mock-muted",e);break;case"color-background-muted":d(o,".mock-main .mock-muted",e);break;case"color-link-foreground":i(o,".mock-main .mock-link",e);break;case"color-link-visited-foreground":i(o,".mock-main .mock-link-visited",e);break;case"color-nav-foreground":i(o,".mock-nav",e),i(o,".mock-nav .mock-link",e);break;case"color-nav-background":d(o,".mock-nav",e);break;case"color-menu-foreground":i(o,".mock-menu",e),i(o,".mock-menu .mock-link",e);break;case"color-menu-background":d(o,".mock-menu",e);break;case"color-menu-selected-foreground":i(o,".mock-menu .mock-link-selected",e);break;case"color-menu-selected-background":d(o,".mock-menu .mock-link-selected",e);break;default:console.error("invalid key ["+n+"]")}}function A(t,n,e){let o=t.querySelectorAll(n);if(o.length==0)throw"empty query selector ["+n+"]";o.forEach(r=>{e(r)})}function d(t,n,e){A(t,n,o=>o.style.backgroundColor=e)}function i(t,n,e){A(t,n,o=>o.style.color=e)}function C(){window.projectforge.Socket=q}var q=class{constructor(n,e,o,r,c){this.debug=n,this.open=e,this.recv=o,this.err=r,this.url=D(c),this.connected=!1,this.pauseSeconds=1,this.pendingMessages=[],this.connect()}connect(){this.connectTime=Date.now(),this.sock=new WebSocket(D(this.url));let n=this;this.sock.onopen=()=>{n.connected=!0,n.pendingMessages.forEach(n.send),n.pendingMessages=[],n.debug&&console.log("WebSocket connected"),n.open("todo")},this.sock.onmessage=e=>{let o=JSON.parse(e.data);n.debug&&console.debug("in",o),n.recv(o)},this.sock.onerror=e=>()=>{n.err("socket",e.type)},this.sock.onclose=()=>{n.connected=!1;let e=n.connectTime?Date.now()-n.connectTime:0;0<e&&e<2e3?(n.pauseSeconds=n.pauseSeconds*2,n.debug&&console.debug(`socket closed immediately, reconnecting in ${n.pauseSeconds} seconds`),setTimeout(()=>{n.connect()},n.pauseSeconds*1e3)):(console.debug("socket closed after ["+e+"ms]"),n.connect())}}disconnect(){}send(n){if(this.debug&&console.debug("out",n),!this.sock)throw"not initialized";if(this.connected){let e=JSON.stringify(n,null,2);this.sock.send(e)}else this.pendingMessages.push(n)}};function D(t){if(t||(t="/connect"),t.indexOf("ws")==0)return t;let n=document.location,e="ws";return n.protocol==="https:"&&(e="wss"),t.indexOf("/")!=0&&(t="/"+t),e+`://${n.host}${t}`}function O(){}function U(){window.projectforge={},y(),v(),b(),L(),M(),H(),S(),N(),C(),O()}document.addEventListener("DOMContentLoaded",U);})();
//# sourceMappingURL=client.js.map
