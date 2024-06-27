<template>
  <div class="container">
    <div class="row">
      <div class="col-12 text-center">
        <h4 class="pt-3">Ubah Kategori</h4>
      </div>
    </div>

    <div class="row">
      <div class="col-3"></div>
      <div class="col-md-6 px-5 px-md-0">
        <form @submit.prevent="editCategory" v-if="category">
          <div class="form-group">
            <label>Nama kategori</label>
            <input type="text" class="form-control" v-model="category.category" required>
          </div>
          <div class="form-group">
            <label for="imageURL">Gambar</label>
            <input type="file" class="form-control-file" id="imageURL" ref="imageURL" @change="handleFileUploadMain">
          </div>
          <button type="submit" class="btn btn-primary">Simpan</button>
        </form>
      </div>
      <div class="col-3"></div>
    </div>
  </div>
</template>

<script>
//import axios from 'axios';

export default {
  data(){
    return {
      category: null,
      imgMain: '',
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
      const slug = this.generateSlug(this.category.category);

      const file = this.$refs.imageURL.files[0];
      if (file) {
        const reader = new FileReader();
        reader.readAsDataURL(file);
        reader.onload = () => {
          let base64 = reader.result.split(',')[1];
          this.imgMain = base64;

          const editCategory = {
            id: this.category.id,
            category: this.category.category,
            image: this.imgMain,
            slug: slug,
          };

          axios.put(`${this.baseURL}api/categories/${this.id}`, editCategory, {
            headers: {
              Authorization: `Bearer ${this.token}`,
            },
          }).then((res) => {
            if (res.data.code === 200) {
              swal({
                text: "Kategori berhasil diubah!",
                icon: "success",
                closeOnClickOutside: false,
              });
              this.$emit("fetchData");
              this.$router.push({ name: 'AdminCategory' });
            }
          }).catch(err => {
            console.error(err);
          });
        };
      } else {
        const editCategory = {
          id: this.category.id,
          category: this.category.category,
          image: this.imgMain,
          slug: slug,
        };

        axios.put(`${this.baseURL}api/categories/${this.id}`, editCategory, {
          headers: {
            Authorization: `Bearer ${this.token}`,
          },
        }).then((res) => {
          if (res.data.code === 200) {
            swal({
              text: "Kategori berhasil diubah!",
              icon: "success",
              closeOnClickOutside: false,
            });
            this.$emit("fetchData");
            this.$router.push({ name: 'AdminCategory' });
          }
        }).catch(err => {
          console.error(err);
        });
      }
    },
    generateSlug(name) {
      const slug = name.toLowerCase().replace(/[^a-z0-9 ]/g, '').replace(/ /g, '-');
      return slug;
    },
    handleFileUploadMain(event) {
      this.imgMain = event.target.files;
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
