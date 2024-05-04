// typescript example:

const btn = document.createElement('button');
let count = 0;
btn.textContent = 'Click me!';

btn.onclick = () => {
  count += 1;
  btn.textContent = `Click me! ${count}`;
};

document.body.appendChild(btn);
