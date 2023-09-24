let arr = [4, 5, 3, 1, 6, 2];
let frames = [];
let n = arr.length;

let frameCounter = 0;
let initialJPosition = 25;

const arrayItems = document.querySelectorAll(".array-item");

const counterJ = document.querySelector(".loop-counter-j");
const counterI = document.querySelector(".loop-counter-i");

const prevBtn = document.querySelector(".prev-btn");
const nextBtn = document.querySelector(".next-btn");

counterJ.style.left = initialJPosition + 'px';

prevBtn.addEventListener("click", () => {
  
});

nextBtn.addEventListener("click", () => {
  frameCounter++;

  const currentFrame = frames[frameCounter];

  counterJ.style.left = (initialJPosition + currentFrame.j * 50) + 'px';

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

  }, 300);
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
          isHighlight: index === j || index === j + 1
        }
      }),
      i: i,
      j: j
    });
  }
}
