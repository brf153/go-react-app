// axios.js

import axios from 'axios';

const instance = axios.create({
  baseURL: 'http://localhost:4000', // Replace with your proxy address
});

export default instance;