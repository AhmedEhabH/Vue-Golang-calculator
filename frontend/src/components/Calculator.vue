<template>
    <div class="hello">

        <div>
            <h1 class="title">Calculator</h1>
        </div>
        <hr>

        <div v-if="errors.length">
            <b>Please correct the following error(s):</b>
            <ul v-for="error in errors" v-bind:key="error">    
                <li class="text-danger">{{ error }}</li>
            </ul>
        </div>

        <div>
            <b-container fluid>

                <b-row class="text-center">

                    <b-col cols="2" class="text-center">
                    </b-col>

                    <b-col cols="4" class="text-center">

                        <form>

                            <div class="field">

                                <label for="" class="label">First Number</label>
                                <b-form-input name="num1" v-model="num1" class="input" type="text" required></b-form-input>

                            </div>

                            <div class="field"  style="margin-top:13px;">

                                <label for="" class="label">Second Number</label>
                                <b-form-input name="num2" v-model="num2" class="input" type="text" required></b-form-input>

                            </div>

                        </form>

                    </b-col>

                    <b-col cols="4" class="text-left">
                        <div><label for="" class="label">Addition: {{ add }}</label></div>
                        <div><label for="" class="label">Multiplication: {{ mul }}</label></div>
                        <div><label for="" class="label">Subtraction: {{ sub }}</label></div>
                        <div><label for="" class="label">Division: {{ div }}</label></div>
                    </b-col>

                </b-row>

            </b-container>
        </div>

        &nbsp;

        <div>
            <b-button variant="primary" v-on:click="postReq()">Calculate</b-button>
        </div>

        &nbsp;
        <hr>
    
     </div>
</template>

<script>
import axios from 'axios';

export default {
    name: 'Calculator',

    data: function(){
        return {
            add: "",
            mul: "",
            sub: "",
            div: "",
            num1: "",
            num2: "",
            errors:[]
        }
    },

    methods: {
        validate: function(){
            if (this.num1 && this.num2)
            {
                let regex=/\b^[0-9]+(.[0-9]+)?$\b/;
                let pat = new RegExp(regex);
                if(pat.exec(this.num1) && pat.exec(this.num2)){
                    if(this.num2 != "0") return true
                }
            }
            return false;
        },
        getErrors(){
            let regex=/\b^[0-9]+(.[0-9]+)?$\b/;
            let pat = new RegExp(regex);
            let errors = []
            if(!this.num1)
            {
                errors.push("Enter number 1");
            }
            if(!this.num2){
                errors.push("Enter number 2");
            }
            if(!pat.exec(this.num1)){
                errors.push("number 1 must be digits");
            }
            if(!pat.exec(this.num2)){
                errors.push("number 2 must be digits");
            }
            if(this.num2 == "0"){
                errors.push("number 2 cant't be zero");
            }
            return errors
        },

        requestSent:function(data){
            axios({
                    method: "POST",
                    url: "http://127.0.0.1:8090/calc",
                    data: data,
                    headers: {
                        "content-type": "text/plain"
                    }
                }).then(result => {
                    this.add = result.data['add'];
                    this.mul = result.data['mul'];
                    this.sub = result.data['sub'];
                    this.div = result.data['div'];

                }).catch(error => {
                    /*eslint-disable*/
                    console.error(error);
                    /*eslint-enable*/
                });
        },

        postReq: function(){
            if (this.validate()) {
                this.errors=[];
                var data = {
                    "num1": parseFloat(this.num1),
                    "num2": parseFloat(this.num2),
                }
                this.requestSent(data);
            } else {
                this.errors = this.getErrors();
            }
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>