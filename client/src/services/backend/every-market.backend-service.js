import axios, { isAxiosError } from 'axios'

class EveryMarketBackendService {
  constructor(baseURL) {
    this.axiosInstance = axios.create({
      baseURL,
      headers: {
        'Content-Type': 'application/json',
      },
    });

    // // Interceptor opcional para agregar token u otra lógica
    // this.axiosInstance.interceptors.request.use((config) => {
    //   const token = localStorage.getItem('token');
    //   if (token) {
    //     config.headers.Authorization = `Bearer ${token}`;
    //   }
    //   return config;
    // });
  }

  async getCategories() {
    try {
      const response = await this.axiosInstance.get(`/product-categories/?page=1&page_size=4`)
      console.log(response.data)
      return response.data
    } catch (error) {
      console.log({ error })
      if (isAxiosError(error)) {
        console.error(`error trying to get products categories`, error.response?.data)
      }
      return []
    }
  }

  async getProductsByCategory(category) {
    try {
      const response = await this.axiosInstance.get(`/products/category/${category}`)
      console.log(response.data)
      return response.data
    } catch (error) {
      if (isAxiosError(error)) {
        console.error(`error trying to get products by category`, error.response.data)
      }
      return []
    }
  }

  async getProducts() {
    return []
  }

  async getProductById(productId) {
    return {}
  }

  createOrder(cartItems, buyer) {
    return {}
  }
}

export const everyMarketBackendService = new EveryMarketBackendService('http://localhost:5000/api')