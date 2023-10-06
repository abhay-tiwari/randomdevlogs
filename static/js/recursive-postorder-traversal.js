let index = -1;
let arr = [];

const onPrevStep = () => {
  if(index === -1) {
    return;
  }

  renderPostorder(false);

  index--;
}

const onNextStep = () => {
  if(index === 12) {
    return;
  }
  
  index++;
  
  renderPostorder(true);
}

const renderPostorder = (isNext) => {
  let node = null;

  switch(index) {
    case 0:
      node = document.querySelector(".one .edge-1");

      if(isNext === true) {
        node.classList.add("traversed-edge");
      }
      else {
        node.classList.remove("traversed-edge");
      }
    break;

    case 1:
      node = document.querySelector(".two .edge-1");

      if(isNext === true) {
        node.classList.add("traversed-edge");
      }
      else {
        node.classList.remove("traversed-edge");
      }

      break;

    case 2:
      node = document.querySelector(".four");

      if(isNext === true) {
        node.classList.add("traversed-node");
        arr.push(4);
      }
      else {
        node.classList.remove("traversed-node");
        arr.pop();
      }
      break;

    case 3:
      node = document.querySelector(".two .edge-2");

      if(isNext === true) {
        node.classList.add("traversed-edge");
      }
      else {
        node.classList.remove("traversed-edge");
      }
      break;

    case 4:
      node = document.querySelector(".five");

      if(isNext === true) {
        node.classList.add("traversed-node");
        arr.push(5);
      }
      else {
        node.classList.remove("traversed-node");
        arr.pop();
      }
      break;

    case 5:
      node = document.querySelector(".two");

      if(isNext === true) {
        node.classList.add("traversed-node");
        arr.push(2);
      }
      else {
        node.classList.remove("traversed-node");
        arr.pop();
      }
      break;
    
    case 6:
      node = document.querySelector(".one .edge-2");

      if(isNext === true) {
        node.classList.add("traversed-edge");
      }
      else {
        node.classList.remove("traversed-edge");
      }
      break;
    
    case 7:
      node = document.querySelector(".three .edge-1");

      if(isNext === true) {
        node.classList.add("traversed-edge");
      }
      else {
        node.classList.remove("traversed-edge");
      }
      break;
    
    case 8:
      node = document.querySelector(".six");

      if(isNext === true) {
        node.classList.add("traversed-node");
        arr.push(6);
      }
      else {
        node.classList.remove("traversed-node");
        arr.pop();
      }
      break;

    case 9:
      node = document.querySelector(".three .edge-2");

      if(isNext === true) {
        node.classList.add("traversed-edge");
      }
      else {
        node.classList.remove("traversed-edge");
      }
      break;
    
    case 10: 
      node = document.querySelector(".seven");

      if(isNext === true) {
        node.classList.add("traversed-node");
        arr.push(7);
      }
      else {
        node.classList.remove("traversed-node");
        arr.pop();
      }
      break;

    case 11:
      node = document.querySelector(".three");

      if(isNext === true) {
        node.classList.add("traversed-node");
        arr.push(3);
      }
      else {
        node.classList.remove("traversed-node");
        arr.pop();
      }
      break;

    case 12:
      node = document.querySelector(".one");

      if(isNext === true) {
        node.classList.add("traversed-node");
        arr.push(1);
      }
      else {
        node.classList.remove("traversed-node");
        arr.pop();
      }
      break;
  }

  renderTraserval();
}

const renderTraserval = () => {
  if(arr.length === 0) {
    document.querySelector(".tree__traversal-output").innerHTML = "";
    return;
  }
 document.querySelector(".tree__traversal-output").innerHTML = `<span>Postorder Traversal of above tree is ${arr.join(', ')}</span>`
}

