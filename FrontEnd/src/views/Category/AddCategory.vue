 <template>
   <div class="container">
     <div class="row">
       <div class="col-12 text-center">
         <h4 class="pt-3">Add new Category</h4>
       </div>
     </div>

     <div class="row">
       <div class="col-3"></div>
       <div class="col-md-6 px-5 px-md-0">
        <form>
          <div class="form-group">
            <label>Category Name</label>
            <input type="text" class="form-control" v-model="categoryName" required>
          </div>
          <div class="form-group">
            <label>ImageURL</label>
            <input type="url" class="form-control" v-model="imageURL" required>
          </div>
          <button type="button" class="btn btn-primary" @click="addCategory">Submit</button>
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
      categoryName : null,
      imageURL : null,
      token: localStorage.getItem('token'),
    }
  },
  props : ["baseURL", "categories"],
  methods : {
    addCategory() {
      const slug = this.generateSlug(this.categoryName);

      const newCategory = {
        category: this.categoryName,
        image: this.imageURL,
        slug: slug,
      }

      axios.post(`${this.baseURL}api/categories`, newCategory, {
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
  mounted(){
    if (!this.token) {
      this.$router.push({ name: 'Signin' });
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
