// HOF (Higer order function)
// yaitu function ber parameter function / function ber return function
//  callback adalah parameter berupa function

// function ber parameter function
// define a function
function getCapture(camera) {
  // invoke the passed function
  camera("bagza");
}

getCapture(function (name) {
  console.log("hai", name);
});

//  function ber return function
// function returnFunc() {
//   return function (name) {
//     return "hai" + " " + name;
//   };
// }

// const sebutnamaku = returnFunc();

// console.log(sebutnamaku("bagza"));
