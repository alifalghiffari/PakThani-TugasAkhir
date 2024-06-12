<template>
  <div class="container">
    <div class="row">
      <div class="col-12 text-center">
        <h4 class="pt-3">Edit Product</h4>
      </div>
    </div>

    <div class="row">
      <div class="col-3"></div>
      <div class="col-md-6 px-5 px-md-0">
        <form @submit.prevent="editProduct" v-if="product">
          <div class="form-group">
            <label>Category</label>
            <select class="form-control" v-model="product.categoryId" required>
              <option v-for="category in categories" :value="category.id">
                {{ category.category }}
              </option>
            </select>
          </div>
          <div class="form-group">
            <label>Name</label>
            <input type="text" class="form-control" v-model="product.name" required>
          </div>
          <div class="form-group">
            <label>Description</label>
            <input type="text" class="form-control" v-model="product.description" required>
          </div>
          <div class="form-group">
            <label>ImageURL</label>
            <input type="url" class="form-control" v-model="product.image" required>
          </div>
          <div class="form-group">
            <label>Price</label>
            <input type="number" class="form-control" v-model.number="product.price" required>
          </div>
          <div class="form-group">
            <label>Stock</label>
            <input type="number" class="form-control" v-model.number="product.quantity" required>
          </div>
          <button type="submit" class="btn btn-primary">Submit</button>
        </form>
      </div>
      <div class="col-3"></div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      product: null,
      categories: [],
      token: localStorage.getItem('token'),
    }
  },
  mounted() {
    if (!this.token) {
      this.$router.push({ name: 'Signin' });
    } else {
      this.id = this.$route.params.id;
      this.product = this.products.find(product => product.id == this.id);
      this.getCategories();
    }
  },
  props: ["baseURL", "products"],
  methods: {
    getCategories() {
      axios.get(`${this.baseURL}api/category`, {
        headers: {
          Authorization: `Bearer ${this.token}`
        }
      }).then((res) => {
        if (res.data.code === 200) {
          this.categories = res.data.data;
          this.$nextTick(() => {
            const category = this.categories.find(c => c.id === this.product.categoryId);
            this.product.categoryId = category ? category.id : null;
          });
        }
      }).catch(err => {
        console.error(err);
      });
    },
    editProduct() {
      const slug = this.generateSlug(this.product.name);

      const editProduct = {
        id: this.product.id,
        category_id: this.product.categoryId,
        name: this.product.name,
        description: this.product.description,
        image: this.product.image,
        price: this.product.price,
        quantity: this.product.quantity,
        slug: slug,
      };

      console.log(editProduct);

      axios.put(`${this.baseURL}/api/products/${this.product.id}`, editProduct, {
        headers: {
          Authorization: `Bearer ${this.token}`,
        },
      }).then((res) => {
        if (res.data.code === 200) {
          swal({
            text: "Product Edit Successfully!",
            icon: "success",
            closeOnClickOutside: false,
          });
          this.$emit("fetchData");
          this.$router.push({ name: 'AdminProduct' });
        }
      }).catch(err => {
        console.error(err);
      });
    },
    generateSlug(name) {
      // Remove special characters and replace spaces with hyphens
      const slug = name.toLowerCase().replace(/[^a-z0-9 ]/g, '').replace(/ /g, '-');
      return slug;
    },
  }
}
</script>

<style scoped>
h4 {
  font-family: 'Roboto', sans-serif;
  color: #484848;
  font-weight: 700;
}
</style>
