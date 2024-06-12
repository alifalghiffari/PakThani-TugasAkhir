<template>
    <div class="container">
        <div class="row">
            <div class="col-12 text-center">
                <h4 class="pt-3">Order Id: {{ orderID }}</h4>
            </div>
        </div>

        <div v-for="(orderItem, index) in orderItems" :key="index" class="row mt-2 pt-3 justify-content-around">
            <div class="col-1"></div>
            <div class="col-md-3 embed-responsive embed-responsive-16by9">
                <img :src="orderItem.product.image" class="w-100 card-img-top embed-responsive-item">
                <hr />
            </div>

            <div class="col-md-5 px-3">
                <div class="card-block px-3">
                    <h6 class="card-title">{{ orderItem.product.name }}</h6>
                    <p id="item-price" class="mb-0 font-weight-bold">${{ orderItem.product.price }} per unit</p>
                    <p id="item-quantity" class="mb-0">Quantity : {{ orderItem.quantity }}</p>
                    <p id="item-total-price" class="mb-0">
                        Total Price : $<span class="font-weight-bold">{{ orderItem.price }}</span>
                    </p>
                </div>
            </div>
            <div class="col-1"></div>
        </div>

        <div class="total-cost pt-2 text-right">
            <h5>Total Cost : $ {{ totalPrice }}</h5>
        </div>
    </div>


</template>

<script>

export default {
    name: 'OrderItems',
    props: ["orderID", "baseURL"],

    data() {
        return {
            orderItems: [],
            order: {},
            token: localStorage.getItem("token"),
            orderID: 0
        }
    },
    computed: {
        totalPrice() {
            return this.orderItems.reduce((total, item) => {
                return total + (item.product.price * item.quantity);
            }, 0);
        }
    },
    methods: {
        getOrder() {
            axios.get(`${this.baseURL}/api/orderitems/admin/${this.orderID}`, {
                headers: {
                    Authorization: `Bearer ${this.token}`
                }
            }).then((res) => {
                if (res.status === 200) {
                    //1 Way Using Order ID
                    this.orderItems = res.data.data
                    //1 Way Using Order
                    // this.order = res.data.data
                    // this.orderItems = this.order.orderItems
                }
                console.log(this.orderItems);
            }).catch((err) => {
                console.error(err);
            });
        }
    },

    mounted() {
        this.orderID = this.$route.params.id;
        this.token = localStorage.getItem("token")
        this.getOrder()
    }
}

</script>

<style scoped>
h4 {
    font-family: 'Roboto', sans-serif;
    color: #484848;
    font-weight: 700;
}

.embed-responsive .card-img-top {
    object-fit: cover;
}
</style>
