let index = -1;
let arr = [];

const onPrevStep = () => {
  if(index === -1) {
    return;
  }

  renderPreorder(false);

  index--;
}

const onNextStep = () => {
  if(index === 12) {
    return;
  }
  
  index++;
  
  renderPreorder(true);
}

const renderPreorder = (isNext) => {
  let node = null;

  switch(index) {
    case 0:
      node = document.querySelector(".one");
      isNext === true ? arr.push(1) : arr.pop();
      isNext === true ? node.classList.add("traversed-node") : node.classList.remove("traversed-node");
    break;

    case 1:
      node = document.querySelector(".one .edge-1");
      isNext === true ? node.classList.add("traversed-edge") : node.classList.remove("traversed-edge"); 
      break;

    case 2:
      node = document.querySelector(".two");
      isNext === true ? arr.push(2) : arr.pop();
      isNext === true ? node.classList.add("traversed-node") : node.classList.remove("traversed-node");
      break;

    case 3:
      node = document.querySelector(".two .edge-1");
      isNext === true ? node.classList.add("traversed-edge") : node.classList.remove("traversed-edge");
      break;

    case 4:
      node = document.querySelector(".four");
      isNext === true ? arr.push(4) : arr.pop();
      isNext === true ? node.classList.add("traversed-node") : node.classList.remove("traversed-node");
      break;

    case 5:
      node = document.querySelector(".two .edge-2");
      isNext === true ? node.classList.add("traversed-edge") : node.classList.remove("traversed-edge");
      break;
    
    case 6:
      node = document.querySelector(".five");
      isNext === true ? arr.push(5) : arr.pop();
      isNext === true ? node.classList.add("traversed-node"): node.classList.remove("traversed-node");
      break;
    
    case 7:
      node = document.querySelector(".one .edge-2");
      isNext === true ? node.classList.add("traversed-edge") : node.classList.remove("traversed-edge");
      break;
    
    case 8:
      node = document.querySelector(".three");
      isNext === true ? arr.push(3) : arr.pop();
      isNext === true ? node.classList.add("traversed-node") : node.classList.remove("traversed-node");
      break;

    case 9:
      node = document.querySelector(".three .edge-1");
      isNext === true ? node.classList.add("traversed-edge") : node.classList.remove("traversed-edge");
      break;
    
    case 10: 
      node = document.querySelector(".six");
      isNext === true ? arr.push(6) : arr.pop();
      isNext === true ? node.classList.add("traversed-node") : node.classList.remove("traversed-node");
      break;

    case 11:
      node = document.querySelector(".three .edge-2");
      isNext === true ? node.classList.add("traversed-edge") : node.classList.remove("traversed-edge");
      break;

    case 12:
      node = document.querySelector(".seven");
      isNext === true ? arr.push(7) : arr.pop();
      isNext === true ? node.classList.add("traversed-node") : node.classList.remove("traversed-node");
      break;
  }

  renderTraserval();
}

const renderTraserval = () => {
  if(arr.length === 0) {
    document.querySelector(".tree__traversal-output").innerHTML = "";
    return;
  }

  document.querySelector(".tree__traversal-output").innerHTML = `<span>Preorder Traversal of above tree is ${arr.join(',')}</span>`
}

