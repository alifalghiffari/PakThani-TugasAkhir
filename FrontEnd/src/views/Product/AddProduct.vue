<template>
  <div class="container">
    <div class="row">
      <div class="col-12 text-center">
        <h4 class="pt-3">Add new Product</h4>
      </div>
    </div>

    <div class="row">
      <div class="col-3"></div>
      <div class="col-md-6 px-5 px-md-0">
        <form @submit.prevent="addProduct">
          <div class="form-group">
            <label>Category</label>
            <select class="form-control" v-model.number="categoryId" required>
              <option v-for="category in categories" :value="category.id">
                {{ category.category }}
              </option>
            </select>
          </div>
          <div class="form-group">
            <label>Name</label>
            <input type="text" class="form-control" v-model="name" required>
          </div>
          <div class="form-group">
            <label>Description</label>
            <input type="text" class="form-control" v-model="description" required>
          </div>
          <div class="form-group">
            <label>ImageURL</label>
            <input type="text" class="form-control" v-model="imageURL" required>
          </div>
          <div class="form-group">
            <label>Price</label>
            <input type="number" class="form-control" v-model.number="price" required>
          </div>
          <div class="form-group">
            <label>Stock</label>
            <input type="number" class="form-control" v-model.number="stock" required>
          </div>
          <button type="submit" class="btn btn-primary">Submit</button>
        </form>
      </div>
      <div class="col-3"></div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      categoryId: null,
      name: null,
      description: null,
      imageURL: null,
      price: null,
      stock: null,
      categories: [],
      token: localStorage.getItem('token'),
    }
  },
  props: ["baseURL"],
  methods: {
    getCategories() {
      axios.get(`${this.baseURL}api/category`, {
        headers: {
          Authorization: `Bearer ${this.token}`
        }
      }).then((res) => {
        if (res.data.code === 200) {
          this.categories = res.data.data;
        }
      }).catch(err => {
        console.error(err);
      });
    },
    addProduct() {
      const slug = this.generateSlug(this.name);

      const newProduct = {
        category_id: this.categoryId,
        name: this.name,
        description: this.description,
        image: this.imageURL,
        price: this.price,
        quantity: this.stock,
        slug: slug,
      };

      axios.post(`${this.baseURL}api/products`, newProduct, {
        headers: {
          Authorization: `Bearer ${this.token}`,
        },
      }).then((res) => {
        if (res.data.code === 200) {
          swal({
            text: "Product Added Successfully!",
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
  },
  mounted() {
    if (!this.token) {
      this.$router.push({ name: 'Signin' });
    } else {
      this.getCategories();
    }
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
