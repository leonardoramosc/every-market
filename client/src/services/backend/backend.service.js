import { everyMarketBackendService } from "./every-market.backend-service";

class BackendService {
  _service;
  constructor(service) {
    this._service = service;
  }

  getCategories() {
    return this._service.getCategories()
  }

  getProductsByCategory(category) {
    return this._service.getProductsByCategory(category);
  }

  getProducts() {
    console.log('GET PRODUCTS')
    return this._service.getProducts();
  }

  async getProductById(productId) {
    return this._service.getProductById(productId);
  }

  createOrder(cartItems, buyer) {
    return this._service.createOrder(cartItems, buyer);
  }
}

export const backendService = new BackendService(everyMarketBackendService);
