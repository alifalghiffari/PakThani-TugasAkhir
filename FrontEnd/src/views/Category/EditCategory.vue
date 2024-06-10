<template>
  <div class="container">
    <div class="row">
      <div class="col-12 text-center">
        <h4 class="pt-3">Edit Category</h4>
      </div>
    </div>

    <div class="row">
      <div class="col-3"></div>
      <div class="col-md-6 px-5 px-md-0">
        <form @submit.prevent="editCategory" v-if="category">
          <div class="form-group">
            <label>Category Name</label>
            <input type="text" class="form-control" v-model="category.category" required>
          </div>
          <div class="form-group">
            <label>ImageURL</label>
            <input type="url" class="form-control" v-model="category.icon" required>
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
  data(){
    return {
      category: null,
      token: localStorage.getItem('token'),
    }
  },
  props : ["baseURL"],
  methods : {
    getCategories(){
      axios.get(`${this.baseURL}api/categories/${this.id}`)
      .then((res) => {
        if(res.data.code == 200) {
          this.category = res.data.data
          console.log(this.category);
        }
      }).catch(err => {
        console.error(err);
      });
    },
    editCategory() {
      //delete this.category["products"]
      const slug = this.generateSlug(this.category.category);

      const editCategory = {
        id: this.category.id,
        category: this.category.category,
        image: this.category.icon,
        slug: slug,
      }

      axios.put(`${this.baseURL}api/categories/${this.id}`, editCategory, {
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
          this.$router.push({ name: 'AdminCategory' });
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
      this.id = this.$route.params.id;
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
