hljs.highlightAll();

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
