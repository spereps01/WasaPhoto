<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			ut: null,
			users: [],
		}
	},
	methods: {
		load() {
			return load
		},
		async getUser() {
	
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/search/" + this.$route.params.username);
				this.users = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},

	},

}
</script>


<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Profilo</h1>
			
		</div>
		<div class="mb-3">
			<a href="javascript:" class="btn btn-primary" @click="getUser()">Info</a>
		</div>
		<LoadingSpinner v-if="loading"></LoadingSpinner>



		<div class="card" v-if="!loading" v-for="u in users">

			<div class="card-body">
				<p class="card-text">
					Username:{{" "  + u.Username }}<br/>
					Numero di foto:{{" "  + u.N_p }}<br/>
					Foto:{{" "  + u.Photos }}<br/>


				</p>
			</div>
		</div>
    </div>
</template>