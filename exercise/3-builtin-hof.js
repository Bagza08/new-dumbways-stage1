// biasanya hof dari javascript biasa digunakan untuk mngolah data yang sifat nya itu array
// ------------------------------------------------------------------>

// forEach --> seperti loop tapi tidak me return data
// const data1 = [1, 2, 3, 4];
// const data1 = [5, 4, 3, 2, 1];

// // data.forEach(function () {});
// // atau bisa ,menggunakan arrow func
// data1.forEach((value, index) => {
//   console.log("nilai :", value);
//   console.log("index :", index);
// });

// ------------------------------------------------------------------>

// map  ---> seperti loop tp dia me return data
// const data2 = [1, 2, 3, 4, 5];

// const double = data2.map(function (value, index) {
//   //   console.log("nilai :", value);
//   //   console.log("index :", index);
//   return value * 2;
// });

// console.log(double);

// ------------------------------------------------------------------>

// filter ---> gunakan ini untuk rating
// const data3 = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

// const filteredData = data3.filter((value, index) => {
//   if (value % 2 == 0) {
//     return true;
//     // return console.log("ini genap :", value);
//   } else {
//     return false;
//   }
// });

// menggunkaan cara simple
// const filteredData = data3.filter((value, index) => value % 2 == 0); //dengan menggnkakan arrow func kita bisa langgsung me return

// console.log(filteredData);

// ------------------------------------------------------------------>

// reduce ---> menjumlahkan, prevvalue , value
// const data4 = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

// const sumData = data4.reduce((prev, value) => prev + value);

// console.log(sumData);
