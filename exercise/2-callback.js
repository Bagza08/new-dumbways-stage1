function hallo() {
  console.log("hallo world");
}

function goodbye() {
  console.log("Sayonara");
}

function print(callback) {
  callback();
}

print(hallo);
print(goodbye);
