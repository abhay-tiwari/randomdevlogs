let front = 0;
let rear = -1;
let maxSize = 6;
let elementCount = 0;

const enqueue = () => {
  const value = document.querySelector(".queue__controls input").value;
 
  hasErrors = false;
 
  if(!value || isNaN(value) === true) {
    hasErrors = true;
    errorMessage = "Please enter a number to enqueue in Queue."
  }

  if( elementCount >= maxSize) {
    hasErrors = true;
    errorMessage = "Queue is full. Use Dequeue to remove element from Queue.";
  }

  if(hasErrors) {
    document.querySelector(".queue__controls-error").innerHTML=`<span color="#FF6B6B">${errorMessage}</span>`;
    document.querySelector(".queue__controls input").value = '';
    return;
  }

  else {
    document.querySelector(".queue__controls-error").innerHTML = ``;
  }

  rear = (rear + 1) % maxSize;
  elementCount += 1;

  switch(rear) {
    case 0:
      document.querySelector(".zero").textContent = value;
      document.querySelector(".queue__rear").style.left = "20px";
    break;

    case 1:
      document.querySelector(".one").textContent = value;
      document.querySelector(".queue__rear").style.left = "70px";
    break;

    case 2:
      document.querySelector(".two").textContent = value;
      document.querySelector(".queue__rear").style.left = "120px";
    break;

    case 3: 
      document.querySelector(".three").textContent = value;
      document.querySelector(".queue__rear").style.left = "170px";
    break;

    case 4:
      document.querySelector(".four").textContent = value;
      document.querySelector(".queue__rear").style.left = "220px";
    break;
    
    case 5: 
      document.querySelector(".five").textContent = value;
      document.querySelector(".queue__rear").style.left = "270px";
    break;
  }

  if(elementCount === 1) {
    document.querySelector(".queue__front").style.left = document.querySelector(".queue__rear").style.left;
  }

  document.querySelector(".queue__controls input").value = '';
}

const dequeue = () => {
  hasErrors = false;

  if( elementCount === 0) {
    hasErrors = true;
    errorMessage = "Queue is Empty. Use Enqueue to add element to Queue."
  }

  if(hasErrors) {
    document.querySelector(".queue__controls-error").innerHTML=`<span color="#FF6B6B">${errorMessage}</span>`;
    return;
  }

  else {
    document.querySelector(".queue__controls-error").innerHTML = ``;
  }

  switch(front) {
    case 0:
      document.querySelector(".zero").textContent = '';
      document.querySelector(".queue__front").style.left = "70px";
    break;

    case 1:
      document.querySelector(".one").textContent = '';
      document.querySelector(".queue__front").style.left = "120px";
    break;

    case 2:
      document.querySelector(".two").textContent = '';
      document.querySelector(".queue__front").style.left = "170px";
    break;

    case 3:
      document.querySelector(".three").textContent = '';
      document.querySelector(".queue__front").style.left = "220px";
    break;

    case 4:
      document.querySelector(".four").textContent = '';
      document.querySelector(".queue__front").style.left = "270px";
    break;

    case 5:
      document.querySelector(".five").textContent = '';
      document.querySelector(".queue__front").style.left = "0";
    break;
  }

  front = (front + 1) % maxSize;
  elementCount -= 1;
}

const handleElInputChange = (e) => {
  if(e.target.value == '') {
    return;
  }

  enqueue();
} 


