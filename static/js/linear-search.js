let initialIState = 20;
let index = 0;

const onPrevButtonClick = () => {
  if(index === 0) {
    return;
  }
  
  else {
    document.querySelector(".linear-search__visualization-data").innerHTML = `<span>arr[i] != 6</span>`;
    document.querySelector(".search-element").style.backgroundColor = "#292F32";
    index--;
  }

  document.querySelector(".loop-counter-i").style.left = initialIState + 50 * index + "px";
}

const onNextButtonClick = () => {
  if(index === 4) {
    return;
  }

  else if(index === 3) {
    index++;
    document.querySelector(".linear-search__visualization-data").innerHTML = `<span>arr[i] == 6</span>`;
    document.querySelector(".search-element").style.backgroundColor = "#FF6B6B"; 
  }

  else {
    index++;
    document.querySelector(".linear-search__visualization-data").innerHTML = `<span>arr[i] != 6</span>`;
  }
  
  document.querySelector(".loop-counter-i").style.left = initialIState + 50 * index + "px";
}
