// OOP (OBJECT ORIENTED PROGRAMING)

// contoh class concept / class adalah cetakan / kelebihan re useable

class Car {
  // properties = warna , jmlh roda, harga
  // method = autoDrive, gas , rem
  //   constructor adalah key untuk menampung nilai
  constructor(color, price) {
    this.color = color;
    this.price = price;
  }

  //   bikin method getinfo sendiri
  getInfo() {
    return `i have car with color ${this.color}, i buy it in ${this.price}`;
  }
}

// ini object
const mobil1 = new Car("red", 20000);
const mobil2 = new Car("blue", 30000);

console.log(mobil1.getInfo());
console.log(mobil2.getInfo());

// dibawah ini adalah contoh perbandingan antara fp dan oop

// FUNCTIONAL PROGRAMING (FP)
// function Car(color, price) {
//   return getInfo(color, price);
// }

// function getInfo(color, price) {
//   return `i have car with color ${color}, i buy it in ${price}`;
// }

// // hukum parameter kayak variabel sesuai urutan
// const x = Car("red", 20000);
// const y = Car("blue", 30000);

// console.log(x);
// console.log(y);
