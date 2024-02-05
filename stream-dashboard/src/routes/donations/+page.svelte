<script lang="ts">
    import {donationsStore} from "$lib/donations";

    donationsStore.subscribe((donations) => {
        console.log(donations);
    });
</script>