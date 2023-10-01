let stack = [];
let maxSize = 6;
let topEl = -1;
let errorMessage = "";
let hasErrors = false;

const handleElInputChange = (e) => {
  if(e.target.value == '') {
    return;
  }

  push();
}

const push = () => {
  let value = document.querySelector(".stack__controls input").value;
  hasErrors = false;
 
  if(!value || isNaN(value) === true) {
    hasErrors = true;
    errorMessage = "Please enter a number to push in Stack."
  }

  if(topEl >= maxSize - 1) {
    hasErrors = true;
    errorMessage = "Stack is full. Use pop to remove element from Stack."
  }

  if(hasErrors) {
    document.querySelector(".stack__controls-error").innerHTML=`<span color="#FF6B6B">${errorMessage}</span>`;
    return;
  }

  topEl += 1;
  
  stack[topEl] = value;
  
  switch(topEl) {
    case 0:
      let stackItem1 = document.querySelector(".zero");
      stackItem1.textContent = value;
      stackItem1.style.opacity = "1";
      stackItem1.style.top = "250px";
      document.querySelector(".stack__top").style.bottom = "15px";
    break;
    case 1:
      let stackItem2 = document.querySelector(".one");
      stackItem2.textContent = value;
      stackItem2.style.opacity = "1";
      stackItem2.style.top = "200px";
      document.querySelector(".stack__top").style.bottom = "65px";
    break;
    case 2:
      let stackItem3 = document.querySelector(".two");
      stackItem3.textContent = value;
      stackItem3.style.top = "150px";
      stackItem3.style.opacity = "1";
      document.querySelector(".stack__top").style.bottom = "115px";
    break;
    case 3:
      let stackItem4 = document.querySelector(".three");
      stackItem4.textContent = value;
      stackItem4.style.top = "100px";
      stackItem4.style.opacity = "1";
      document.querySelector(".stack__top").style.bottom = "165px";
    break;
    case 4:
      let stackItem5 = document.querySelector(".four");
      stackItem5.textContent = value;
      stackItem5.style.top = "50px";
      stackItem5.style.opacity = "1";
      document.querySelector(".stack__top").style.bottom = "215px";
    break;
    case 5:
      let stackItem6 = document.querySelector(".five");
      stackItem6.textContent = value;
      stackItem6.style.top = "0px";
      stackItem6.style.opacity = "1";
      document.querySelector(".stack__top").style.bottom = "265px";
    break;
  }
  
  document.querySelector(".stack__controls input").value = '';
}

const pop = () => {
  let x = stack[topEl];
  topEl -= 1;
  debugger;

  switch(topEl) {
    case -1:
      let stackItem1 = document.querySelector(".zero");
      setTimeout(() => {
        stackItem1.style.opacity = "0";
      }, 300);
      stackItem1.style.top = "-70px";
      document.querySelector(".stack__top").style.bottom = "-11px";
    break;

    case 0:
      let stackItem2 = document.querySelector(".one");
      setTimeout(() => {
        stackItem2.style.opacity = "0";
      }, 300);
      stackItem2.style.top = "-70px";
      document.querySelector(".stack__top").style.bottom = "15px";
    break;

    case 1:
      let stackItem3 = document.querySelector(".two");
      stackItem3.style.top = "-70px";
      setTimeout(() => {
        stackItem3.style.opacity = "0";
      }, 300);
      document.querySelector(".stack__top").style.bottom = "65px";
    break;

    case 2:
      let stackItem4 = document.querySelector(".three");
      stackItem4.style.top = "-70px";
      setTimeout(() => {
        stackItem4.style.opacity = "0";
      }, 300);
      document.querySelector(".stack__top").style.bottom = "115px";
    break;

    case 3:
      let stackItem5 = document.querySelector(".four");
      stackItem5.style.top = "-70px";
      setTimeout(() => {
        stackItem5.style.opacity = "0";
      }, 300);      
      document.querySelector(".stack__top").style.bottom = "165px";
    break;

    case 4:
      let stackItem6 = document.querySelector(".five");
      stackItem6.style.top = "-70px";
      setTimeout(() => {
        stackItem6.style.opacity = "0";
      }, 300);      
      document.querySelector(".stack__top").style.bottom = "215px";
    break;
  }
}
