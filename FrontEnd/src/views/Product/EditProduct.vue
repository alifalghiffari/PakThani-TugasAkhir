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
            <select class="form-control" v-model.number="product.categoryId" required>
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
            <label for="imageURL">Main Image</label>
            <input type="file" class="form-control-file" id="imageURL" ref="imageURL" @change="handleFileUploadMain">
          </div>
          <div class="custom-file mb-3">
            <input multiple type="file" class="custom-file-input" id="image1" name="image1" ref="image1" @change="handleFileUpload">
            <label class="custom-file-label" for="image1">Image Optional</label>
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
import axios from 'axios';

export default {
  data() {
    return {
      product: null,
      categories: [],
      token: localStorage.getItem('token'),
      base64Images: [],
      imgMain: ''
    }
  },
  props: ["baseURL", "products"],
  mounted() {
    if (!this.token) {
      this.$router.push({ name: 'Signin' });
    } else {
      this.id = this.$route.params.id;
      this.product = this.products.find(product => product.id == this.id);
      this.getCategories();
    }
  },
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

      const img = this.$refs.imageURL.files[0];
      if (img) {
        const reader = new FileReader();
        reader.readAsDataURL(img);
        reader.onload = () => {
          let base64 = reader.result.split(',')[1];
          this.imgMain = base64;
          this.updateProduct(slug, this.imgMain);
        };
      } else {
        this.updateProduct(slug, null);
      }
    },
    updateProduct(slug, imgMain) {
      const base64Images = [];
      if (this.$refs.image1.files.length > 0) {
        Array.from(this.$refs.image1.files).forEach(file => {
          const reader = new FileReader();
          reader.readAsDataURL(file);
          reader.onload = () => {
            let base64String = reader.result.split(',')[1];
            base64Images.push(base64String);
            if (base64Images.length === this.$refs.image1.files.length) {
              this.submitProductUpdate(slug, imgMain, base64Images);
            }
          };
        });
      } else {
        this.submitProductUpdate(slug, imgMain, base64Images);
      }
    },
    submitProductUpdate(slug, imgMain, base64Images) {
      const editProduct = {
        id: this.product.id,
        category_id: this.product.categoryId,
        name: this.product.name,
        description: this.product.description,
        price: this.product.price,
        quantity: this.product.quantity,
        slug: slug
      };

      if (imgMain) {
        editProduct.image = imgMain;
      }

      if (base64Images.length > 0 ) {
        editProduct.images = base64Images;
      }
      console.log(editProduct);
    
      axios.put(`${this.baseURL}api/products/${this.product.id}`, editProduct, {
        headers: {
          Authorization: `Bearer ${this.token}`
        }
      }).then((res) => {
        if (res.data.code === 200) {
          swal({
            text: "Product Edit Successfully!",
            icon: "success",
            closeOnClickOutside: false
          });
          this.$emit("fetchData");
          this.$router.push({ name: 'AdminProduct' });
        }
      }).catch(err => {
        console.error(err);
      });
    },
    generateSlug(name) {
      const slug = name.toLowerCase().replace(/[^a-z0-9 ]/g, '').replace(/ /g, '-');
      return slug;
    },
    handleFileUploadMain(event) {
      this.imgMain = event.target.files;
    },
    handleFileUpload(event) {
      const files = event.target.files;
      const base64Images = [];
      Array.from(files).forEach(file => {
        const reader = new FileReader();
        reader.readAsDataURL(file);
        reader.onload = () => {
          let base64String = reader.result.split(',')[1];
          base64Images.push(base64String);
          if (base64Images.length === files.length) {
            this.base64Images = base64Images;
          }
        };
      });
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
