// kegunaan polymorphism adalah menimpa mthod parent dengan membuat method child

class Vehicle {
  drive() {
    return "The vehicle is driving";
  }
}

class Car extends Vehicle {
  drive() {
    return "the car is driving";
  }
}

class ElectricCar extends Vehicle {
  drive() {
    return "the electric car is driving";
  }
}

const myVehicle = new Vehicle();
const myCar = new Car();
const myElectricCar = new ElectricCar();

console.log(myVehicle.drive());
console.log(myCar.drive());
console.log(myElectricCar.drive());
