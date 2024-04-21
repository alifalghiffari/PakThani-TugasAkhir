<template>
  <div class="adminPanel container">
    <div class="row">
      <div class="col">
        <router-link :to="{name : 'AdminProduct'}">
          <button type="button" class="btn btn-primary btn-lg">Admin Products</button>
        </router-link>
      </div>
      <div class="col">
        <router-link :to="{name : 'AdminCategory'}">
          <button type="button" class="btn btn-primary btn-lg">Admin Categories</button>
        </router-link>
      </div>
      <div class="col">
        <router-link :to="{name : 'Gallery'}">
          <button type="button" class="btn btn-primary btn-lg">Admin Gallery</button>
        </router-link>
      </div>
      <div class="col">
        <router-link :to="{name : 'Order'}">
          <button type="button" class="btn btn-primary btn-lg">Admin Order</button>
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  mounted() {
    if (!localStorage.getItem('token')) {
      this.$router.push({name : 'Signin'});
    }
  }
}
</script>


<style>
.adminPanel{
  display : flex;
  flex-direction: column;
  align-items: center;
}
.adminPanel button{
  font-size: 30px;
  width: fit-content;
  margin: 55px 0;
}
</style>