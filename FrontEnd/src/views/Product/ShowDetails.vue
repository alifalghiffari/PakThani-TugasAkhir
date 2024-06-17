<template>
  <div class="container">
    <div class="row pt-5">
      <div class="col-md-1"></div>
      <div class="col-md-4 col-12">
        <div id="carouselExampleControls" class="carousel slide" data-ride="carousel">
          <div class="carousel-inner">
            <div class="carousel-item active">
              <img class="d-block w-100" :src="getImagePathMain(product.image)" alt="Product Image">
            </div>
            <div class="carousel-item" v-for="img in product.images" :key="img.id">
              <img class="d-block w-100" :src="getImagePath(img.image)" alt="Product Image">
            </div>
          </div>
          <a class="carousel-control-prev" href="#carouselExampleControls" role="button" data-slide="prev">
            <span class="carousel-control-prev-icon" aria-hidden="true"></span>
            <span class="sr-only">Previous</span>
          </a>
          <a class="carousel-control-next" href="#carouselExampleControls" role="button" data-slide="next">
            <span class="carousel-control-next-icon" aria-hidden="true"></span>
            <span class="sr-only">Next</span>
          </a>
        </div>
      </div>
      <div class="col-md-6 col-12 pt-3 pt-md-0">
        <h4>{{ product.name }}</h4>
        <h6 class="category font-italic">{{ category.category }}</h6>
        <h6 class="font-weight-bold">Rp. {{ product.price }}</h6>
        <p>
          {{ product.description }}
        </p>

        <div class="d-flex flex-row justify-content-between">
          <div class="quantity-container ">
            <span class="quantity-label">Quantity</span>
            <div class="quantity-input-container">
              <button class="quantity-btn" @click="decrementQuantity" :disabled="quantity <= 1">-</button>
              <input class="form-control num text-center" type="number" v-model.number="quantity" min="1" />
              <button class="quantity-btn" @click="incrementQuantity">+</button>
            </div>
          </div>

          <div class="input-group col-md-3 col-4 p-0">
            <button type="button" id="add-to-cart-button" class="btn" @click="addToCart(this.id)">
              Add to Cart
              <ion-icon name="cart-outline" v-pre></ion-icon>
            </button>
          </div>
        </div>
        <div class="pt-2">
          <p>Stock: <b>{{ product.quantity }}</b></p>
        </div>

        <div class="features pt-3">
          <h5><strong>Features</strong></h5>
          <ul>
            <li>Lorem ipsum dolor sit amet consectetur adipisicing elit.</li>
            <li>Officia quas, officiis eius magni error magnam voluptatem</li>
            <li>nesciunt quod! Earum voluptatibus quaerat dolorem doloribus</li>
            <li>molestias ipsum ab, ipsa consectetur laboriosam soluta et</li>
            <li>ut doloremque dolore corrupti, architecto iusto beatae.</li>
          </ul>
        </div>
      </div>
      <div class="col-md-1"></div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      product: {},
      category: [],
      id: null,
      token: localStorage.getItem('token'),
      isAddedToWishlist: false,
      wishlistString: "Add to wishlist",
      quantity: 1,
    };
  },
  props: ["baseURL", "products", "category"],
  methods: {
    getImagePath(image) {
      try {
        return require(`../../assets/img/${image}`);
      } catch (e) {
        return ''; 
      }
    },
    getImagePathMain(image) {
      try {
        return require(`../../assets/img-main/${image}`);
      } catch (e) {
        return ''; 
      }
    },
    incrementQuantity() {
      this.quantity++
    },
    decrementQuantity() {
      if (this.quantity > 1) {
        this.quantity--
      }
    },
    addToCart(productId) {
      productId = parseInt(productId, 10);
      console.log(this.product);

      if (!this.token || this.quantity === 0) {
        swal({
          text: "Please log in first! or check your quantity",
          icon: "error",
        });
        return;
      } else if (this.quantity > this.product.quantity) {
        swal({
          text: "Product out of stock",
          icon: "error",
        });
        return;
      } else {
        axios
          .post(`${this.baseURL}api/carts`, {
            product_id: productId,
            quantity: this.quantity,
          }, {
            headers: {
              Authorization: `Bearer ${this.token}`,
            },
          })
          .then((response) => {
            if (response.data.code === 200) {
              swal({
                text: "Product Added to the cart!",
                icon: "success",
                closeOnClickOutside: false,
              });
              this.$emit("fetchData");
            }
          })
          .catch((error) => {
            console.log(error);
          });
      }
    },
  },
  mounted() {
    this.id = this.$route.params.id;
    this.product = this.products.find((product) => product.id == this.id);
    this.category = this.category.find(
      (category) => category.id == this.product.categoryId
    );
    this.token = localStorage.getItem("token");
  },
};
</script>

<style>
.category {
  font-weight: 400;
}

input::-webkit-outer-spin-button,
input::-webkit-inner-spin-button {
  -webkit-appearance: none;
  margin: 0;
}

input[type="number"] {
  -moz-appearance: textfield;
}

#add-to-cart-button {
  background-color: #febd69;
}

#wishlist-button {
  background-color: #b9b9b9;
  border-radius: 0;
}

#show-cart-button {
  background-color: #131921;
  color: white;
  border-radius: 0;
}

.quantity-container {
  display: flex;
  align-items: center;
}

.quantity-label {
  margin-right: 10px;
}

.quantity-input-container {
  display: flex;
  align-items: center;
}

.quantity-btn {
  padding: 5px 10px;
  border: 1px solid #ccc;
  background-color: #f5f5f5;
  cursor: pointer;
}

.quantity-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.num {
  width: 80px;
  text-align: center;
  padding: 5px;
  border: 1px solid #ccc;
}
</style>
