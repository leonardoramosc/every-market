class EveryMarketBackendService {

  async getProductsByCategory(category) {
    return []
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

export const everyMarketBackendService = new EveryMarketBackendService()