<template>
    <div class="">
      <div class="col-12 justify-content-center d-flex flex-row pt-5">
        <div id="div_class" class="flex-item border">
          <h2 class="pt-4 pl-4">Input Your Address</h2>
          <form v-if="hasAddress" @submit.prevent="addAddress" id="addAdress" class="pt-4 pl-4 pr-4">
            <div class="form-group">
              <label>Kabupaten</label>
              <input type="text" name="kabupaten" id="kabupaten" class="form-control" v-model="kabupaten" required/>
            </div>
            <div class="form-row">
              <div class="col">
                <div class="form-group">
                  <label>Kecamatan</label>
                  <input type="text" name="kecamatan" id="kecmaatan" class="form-control" v-model="kecamatan" required/>
                </div>
              </div>
              <div class="col">
                <div class="form-group">
                  <label>Kelurahan</label>
                  <input type="text" name="kelurahan" id="kelurahan" class="form-control" v-model="kelurahan" required/>
                </div>
              </div>
            </div>
            <div class="form-group">
              <label>Alamat</label>
              <input
                type="text" name="alamat" id="alamat" class="form-control" v-model="alamat" required/>
            </div>
            <div class="form-group">
              <label>Note</label>
              <textarea name="note" class="form-control" id="note" v-model="note" cols="30" rows="4"></textarea>
            </div>
            <div class="form-group">
              <input type="text" name="id" id="id" value="" class="form-control" hidden/>
            </div>
            <div class="form-group">
              <input type="text" name="userId" id="userId" value="" class="form-control" hidden/>
            </div>
            <div class="form-row">
              <div class="col">
                <p class="text-center">
                  <router-link
                    class="btn btn-dark text-center mx-auto px-5 py-1 mb-2" :to="{ name: 'Cart' }">Back To Cart</router-link>
                </p>
              </div>
              <div class="col">
                <p class="text-center">
                  <button class="btn btn-dark text-center mx-auto px-5 py-1 mb-2" type="submit">Update</button>
                </p>
              </div>
            </div>
          </form>
          <form v-else @submit.prevent="updateAddress" id="updateAdress" class="pt-4 pl-4 pr-4">
            <div class="form-group">
              <label>Kabupaten</label>
              <input type="text" name="kabupaten" id="kabupaten" class="form-control" v-model="kabupaten" required/>
            </div>
            <div class="form-row">
              <div class="col">
                <div class="form-group">
                  <label>Kecamatan</label>
                  <input type="text" name="kecamatan" id="kecmaatan" class="form-control" v-model="kecamatan" required/>
                </div>
              </div>
              <div class="col">
                <div class="form-group">
                  <label>Kelurahan</label>
                  <input type="text" name="kelurahan" id="kelurahan" class="form-control" v-model="kelurahan" required/>
                </div>
              </div>
            </div>
            <div class="form-group">
              <label>Alamat</label>
              <input
                type="text" name="alamat" id="alamat" class="form-control" v-model="alamat" required/>
            </div>
            <div class="form-group">
              <label>Note</label>
              <textarea name="note" class="form-control" id="note" v-model="note" cols="30" rows="4"></textarea>
            </div>
            <div class="form-group">
              <input type="text" name="id" id="id" value="" class="form-control" hidden/>
            </div>
            <div class="form-group">
              <input type="text" name="userId" id="userId" value="" class="form-control" hidden/>
            </div>
            <div class="form-row">
              <div class="col">
                <p class="text-center">
                  <router-link
                    class="btn btn-dark text-center mx-auto px-5 py-1 mb-2" :to="{ name: 'Cart' }">Back To Cart</router-link>
                </p>
              </div>
              <div class="col">
                <p class="text-center">
                  <button class="btn btn-dark text-center mx-auto px-5 py-1 mb-2" type="submit">Update</button>
                </p>
              </div>
            </div>
          </form>
        </div>
      </div>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                id: null,
                kabupaten: null,
                kecamatan: null,
                kelurahan: null,
                alamat: null,
                note: null,
                token: localStorage.getItem('token'),
                address: []
            };
        },
        props: ["baseURL", "address"],
        computed: {
          hasAddress() {
            return this.id === null;
          }
        },
        methods: {
            fetchAddress() {
              axios.get(`${this.baseURL}api/address/users`, {
                headers: {
                  Authorization: `Bearer ${this.token}`,
                }
              }).then((response) => {
                if (response.data.code === 200 && response.data.data.length > 0) {
                  const result = response.data.data[0];
                  this.id = result.id;
                  this.kabupaten = result.kabupaten;
                  this.kecamatan = result.kecamatan;
                  this.kelurahan = result.kelurahan;
                  this.alamat = result.alamat;
                  this.note = result.note;
                  console.log(result);
                } else {
                  this.id = null; 
                }
              }).catch((err) => {
                console.log(err);
              })
            },
            addAddress() {
                axios.post(`${this.baseURL}api/address`, {
                    kabupaten: this.kabupaten,
                    kecamatan: this.kecamatan,
                    kelurahan: this.kelurahan,
                    alamat: this.alamat,
                    note: this.note,
                }, {
                    headers: {
                        Authorization: `Bearer ${this.token}`,
                    },
                }).then((response) => {
                    if (response.data.code === 200) {
                        this.fetchAddress();
                        console.log(response.data.data);
                        this.$emit("fecthData");
                        swal("Success", "Address added successfully", "success");
                    }
                }).catch((error) => {
                    console.log(error);
                })
            },
            updateAddress() {
              console.log(this.id);
              axios.put(`${this.baseURL}api/address/update`, {
                id: this.id,
                kabupaten: this.kabupaten,
                kecamatan: this.kecamatan,
                kelurahan: this.kelurahan,
                alamat: this.alamat,
                note: this.note,
              }, {
                headers: {
                  Authorization: `Bearer ${this.token}`,
                }
              }).then((response) => {
                if (response.data.code === 200) {
                  console.log(response.data.data);
                  this.fetchAddress();
                  swal("Success", "Address updated successfully", "success");
                }
              }).catch((err) => {
                console.log(err);
              })
            }
        },
        mounted() {
            this.fetchAddress();
            this.token = localStorage.getItem("token");
        }
    }
</script>

<style>
@media only screen and (min-width: 992px) {
  #div_class {
    width: 40%;
  }
}

</style>