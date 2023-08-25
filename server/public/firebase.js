import { initializeApp } from "https://www.gstatic.com/firebasejs/10.1.0/firebase-app.js";

import {
  getAuth,
  createUserWithEmailAndPassword,
  signInWithEmailAndPassword,
} from "https://www.gstatic.com/firebasejs/10.1.0/firebase-auth.js";

const apiKey = "AIzaSyCYMDy_Rf1o82MmWVyI0p0fRwyJ2w2dcTs";
const authDomain = "cadavre-exquis-9c7af.firebaseapp.com";

const firebaseConfig = {
  apiKey,
  authDomain,
}

const app = initializeApp(firebaseConfig);

const auth = getAuth(app);

function signIn(email, password) {
  signInWithEmailAndPassword(auth, email, password)
    .then((userCredential) => {
      const user = userCredential.user;
      document.cookie = `accessToken=${user.accessToken}`;

      htmx.ajax("GET", '/home', '#main')
    })
    .catch((error) => {
      const errorCode = error.code;
      const errorMessage = error.message;
      // htmx.ajax("GET", '/home', '#main') should show error message
    });
}

function signUp(email, password) {
  createUserWithEmailAndPassword(auth, email, password)
    .then((userCredential) => {
      const user = userCredential.user;
      document.cookie = `accessToken=${user.accessToken}`;
      htmx.ajax("GET", '/home', '#main')
    })
    .catch((error) => {
      const errorCode = error.code;
      const errorMessage = error.message;
      // htmx.ajax("GET", '/home', '#main') should show error message
    });
}

window.signIn = signIn
window.signUp = signUp