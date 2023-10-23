hljs.highlightAll();

if(document.querySelectorAll(".copy-btn").length <= 0) {
  let preElements = document.querySelectorAll("pre");

  preElements.forEach(x => {
    x.style.position = "relative";

    let button = document.createElement("button");
    button.classList.add("btn");
    button.classList.add("copy-btn");
    button.classList.add("copy-button");

    let copyIcon = document.createElement("i");
    copyIcon.classList.add("fa-regular");
    copyIcon.classList.add("fa-copy");

    let checkIcon = document.createElement("i");
    checkIcon.classList.add("fa-solid");
    checkIcon.classList.add("fa-check");

    button.appendChild(copyIcon);
    button.appendChild(checkIcon);

    button.onclick = function() {
      let codeBlock = x.querySelector("code");
      const range = document.createRange();
  
      range.selectNode(codeBlock);

      window.getSelection().removeAllRanges();
      window.getSelection().addRange(range);

      if(!navigator.clipboard) {
        document.execCommand(range.toString());
      }

      else {
        navigator.clipboard.writeText(range.toString());
      }

      window.getSelection().removeAllRanges();
 
      button.classList.add("copied");

      setTimeout(() => {
        button.classList.remove("copied");
      }, 1000);
    }
    x.appendChild(button);
  });
}


const onCopyButtonClick = (element) => {
  let codeBlock = element.parentElement.querySelector("code");

  const range = document.createRange();
  
  range.selectNode(codeBlock);

  window.getSelection().removeAllRanges();
  window.getSelection().addRange(range);

  if(!navigator.clipboard) {
    document.execCommand(range.toString());
  }

  else {
    navigator.clipboard.writeText(range.toString());
  }

  window.getSelection().removeAllRanges();
 
  element.classList.add("copied");

  setTimeout(() => {
    element.classList.remove("copied");
  }, 1000);

}
