<template>
  <div class="container">
    <!--    Logo Div-->
    <div class="row">
      <div class="col-12 text-center pt-3">
        <router-link :to="{ name: 'Home' }">
          <img id="logo" src="../assets/logo.png" />
        </router-link>
      </div>
    </div>

    <div class="row">
      <div class="col-12 justify-content-center d-flex flex-row pt-5">
        <div id="signup-div" class="flex-item border">
          <h2 class="pt-4 pl-4">Create Account</h2>
          <form @submit="signup" class="pt-4 pl-4 pr-4">
            <div class="form-group">
              <label>Email</label>
              <input
                type="email"
                name="email"
                id="email"
                class="form-control"
                v-model="email"
                required
              />
            </div>
            <div class="form-row">
              <div class="col">
                <div class="form-group">
                  <label>Name</label>
                  <input
                    type="name"
                    name="username"
                    id="username"
                    class="form-control"
                    v-model="username"
                    required
                  />
                </div>
              </div>
              <div class="col">
                <div class="form-group">
                  <label>No HP</label>
                  <input
                    type="number"
                    name="nohp"
                    id="nohp"
                    class="form-control w-100"
                    v-model.number="nohp"
                    required
                  />
                </div>
              </div>
            </div>
            <div class="form-group">
              <label>Password</label>
              <input
                type="password"
                name="password"
                id="password"
                class="form-control"
                v-model="password"
                required
                :class="{ 'is-invalid': password && !passwordRegex.test(password) }"
              />
              <div class="invalid-feedback" v-if="password && !passwordRegex.test(password)">
                Password must be at least 8 characters long, contain at least one uppercase letter, one lowercase letter, one digit, and one special character.
              </div>
            </div>
            <div class="form-group">
              <label>Confirm Password</label>
              <input
                type="password"
                class="form-control"
                v-model="passwordConfirm"
                required
              />
            </div>
            <!-- <div class="form-group">
              <input 
              type="text"
              name="role"
              id="role"
              class="form-control"
              
              hidden
              />
            </div> -->
            <button type="submit" class="btn btn-primary mt-2 py-0">
              Create Account
            </button>
          </form>
          <hr />
          <small class="form-text text-muted pt-2 pl-4 text-center"
            >Sudah Punya Akun?</small
          >
          <p class="text-center">
            <router-link
              class="btn btn-dark text-center mx-auto px-5 py-1 mb-2" type="submit"
              :to="{ name: 'Signin' }"
              >Log In Disini</router-link
            >
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: "Signup",
    props: ["baseURL"],
    data() {
      return {
        email: null,
        username: null,
        nohp: null,
        password: null,
        passwordConfirm: null,
        role: true,
        passwordRegex: /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$/
      };
    },
    methods: {
      async signup(e) {
        e.preventDefault();
        // if the password matches and meets the requirements
        if (this.password === this.passwordConfirm && this.passwordRegex.test(this.password)) {
          // make the post body
          const user = {
            email: this.email,
            username: this.username,
            no_telepon: this.nohp,
            password: this.password,
            role: this.role,
          };
          console.log(user);

          // call the API
          await axios
            .post(`http://localhost:3000/api/users/register`, user)
            .then((res) => {
              console.log(res);
              // redirect to home page
              this.$router.replace("/signin");
              swal({
                text: "User signup successful. Please Login",
                icon: "success",
                closeOnClickOutside: false,
              });
            })
            .catch((err) => {
              console.log(err);
            });
        } else {
          // passwords are not matching or do not meet the requirements
          swal({
            text: "Error! Passwords are not matching or do not meet the requirements",
            icon: "error",
            closeOnClickOutside: false,
          });
        }
      },
    },
  };
</script>

<style scoped>
.btn-dark {
  background-color: #e7e9ec;
  color: #000;
  font-size: smaller;
  border-radius: 0;
  border-color: #adb1b8 #a2a6ac #a2a6ac;
}

.btn-primary {
  background-color: #f0c14b;
  color: black;
  border-color: #a88734 #9c7e31 #846a29;
  border-radius: 0;
}

#logo {
  width: 150px;
}

@media only screen and (min-width: 992px) {
  #signup-div {
    width: 40%;
  }
}

</style>
