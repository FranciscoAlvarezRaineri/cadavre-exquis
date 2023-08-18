import { initializeApp } from "https://www.gstatic.com/firebasejs/10.1.0/firebase-app.js";

// Add Firebase products that you want to use
import {
  getAuth,
  createUserWithEmailAndPassword,
  signInWithEmailAndPassword,
} from "https://www.gstatic.com/firebasejs/10.1.0/firebase-auth.js";

// TODO: Replace the following with your app's Firebase project configuration
const firebaseConfig = {
  apiKey: "AIzaSyCYMDy_Rf1o82MmWVyI0p0fRwyJ2w2dcTs",
  authDomain: "cadavre-exquis-9c7af.firebaseapp.com",
  projectId: "cadavre-exquis-9c7af",
  storageBucket: "cadavre-exquis-9c7af.appspot.com",
  messagingSenderId: "950495760420",
  appId: "1:950495760420:web:57784044dd9be8f95548b1",
  measurementId: "G-80DH0CRH6R",
};

const app = initializeApp(firebaseConfig);

const auth = getAuth();

/*       createUserWithEmailAndPassword(auth, email, password)
  .then((userCredential) => {
    // Signed in
    const user = userCredential.user;
    // ...
  })
  .catch((error) => {
    const errorCode = error.code;
    const errorMessage = error.message;
    // ..
  }); */

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

window.signIn = signIn