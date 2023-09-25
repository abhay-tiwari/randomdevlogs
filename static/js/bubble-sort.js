let arr = [4, 5, 3, 1, 6, 2];
let frames = [];
let n = arr.length;

let frameCounter = 0;
const initialJPosition = 25;
const initialIPosotion = 260;
const arrayItems = document.querySelectorAll(".array-item");

const counterJ = document.querySelector(".loop-counter-j");
const counterI = document.querySelector(".loop-counter-i");

const prevBtn = document.querySelector(".prev-btn");
const nextBtn = document.querySelector(".next-btn");
const copyBtn = document.querySelector(".copy-btn");

counterJ.style.left = initialJPosition + 'px';

const renderFrame = (currentFrame) => {
  counterJ.style.left = (initialJPosition + currentFrame.j * 50) + 'px';
  counterI.style.left = (initialIPosotion - currentFrame.i * 50) + 'px';

  arrayItems.forEach((x, index) => {
  
    let el = frames[0].elements[index].value;
    let elIndex = currentFrame.elements.findIndex(y => y.value === el);
    
    if(elIndex !== -1 && currentFrame.elements[elIndex].isHighlight === true) {
      x.style.backgroundColor = "#FF6B6B";
    }
    else {
      x.style.backgroundColor = "#292F36";
    }

  });

  setTimeout(() => {
    arrayItems.forEach((x, index) => {
    let el = frames[0].elements[index].value;
    let elIndex = currentFrame.elements.findIndex(y => y.value === el);
    if(elIndex !== -1) {
      x.style.left = elIndex * 50 + 'px';
    }
  });

  }, 500);

}

prevBtn.addEventListener("click", () => {
  frameCounter--;
  
  if(frameCounter === -1) {
    return;
  }

  renderFrame(frames[frameCounter]);

});

nextBtn.addEventListener("click", () => {
  frameCounter++;

  if(frameCounter >= frames.length) {
    return;
  }

  renderFrame(frames[frameCounter]);
});

copyBtn.addEventListener("click", (event) => {
  let code = document.querySelector("code");
  
  const range = document.createRange();
  range.selectNode(code);
  window.getSelection().removeAllRanges();
  window.getSelection().addRange(range);

  if(!navigator.clipboard) {
    document.execCommand(range.toString());
  }

  else {
    navigator.clipboard.writeText(range.toString());
  }

  window.getSelection().removeAllRanges();
 
  document.querySelector(".copy-btn").classList.add("copied");

  setTimeout(() => {
    document.querySelector(".copy-btn").classList.remove("copied");
  }, 1000);
});


for(let i=0; i< n - 1; i++) {
  for(let j=0; j < n - i - 1; j++) {
    let isSwapped = false;

    if(arr[j] > arr[j + 1]) {
      let temp = arr[j];
      arr[j] = arr[j + 1];
      arr[j + 1] = temp;
      isSwapped = true;
   }

    frames.push({
      elements: arr.map((x, index) => {
        return {
          value: x,
          isHighlight: (index === j || index === j + 1) && isSwapped
        }
      }),
      i: i,
      j: j
    });
  }
}
