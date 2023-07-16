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
    if (this.color == "") {
      return `i have car, i buy it in ${this.price}`;
    }
    return `i have car with color ${this.color}, i buy it in ${this.price}`;
  }
}

// bikin class baru yaitu ElectricCar yang class turunan dari class car
class ElectricCar extends Car {
  constructor(color, price, batteryCapacity) {
    // kita meminjam kekuatan orang tua yaitu class car menggunakan super
    super(color, price);
    this.batteryCapacity = batteryCapacity;
  }

  //   polymorphism concepts --> menimpa
  //   kita bisa replace method dari parent di atas dengfan bikin method baru
  getInfo() {
    return `i have car with color ${this.color}, i buy it in ${this.price}, with battery capacity ${this.batteryCapacity}`;
  }
}

//   ini object
const myElectricCar = new ElectricCar("red", 20000, 200);

// const mobil1 = new Car("red", 20000);
// const mobil2 = new Car("blue", 30000);

console.log(myElectricCar.getInfo());
// console.log(mobil2.getInfo());
