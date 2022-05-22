import axios from 'axios';

const axiosInstance = axios.create();

axiosInstance.defaults.baseURL = 'http://localhost:8080';

// axiosInstance.defaults.baseURL = '/';

export default axiosInstance;
