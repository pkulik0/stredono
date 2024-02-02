<script lang="ts">
    import axios from 'axios';
    import {SendDonateRequest} from "$lib/functions_pb";
    import {PUBLIC_SENT_DONATE_LINK} from "$env/static/public";

    const donate = async () => {
        const sdReq = new SendDonateRequest({
            amount: "5.37",
            currency: "PLN",
            email: "abc@ede.com",
            sender: "abc",
            recipient: "123",
            message: "Test"
        });

        const res = await axios.post(PUBLIC_SENT_DONATE_LINK, sdReq.toBinary())
        console.log(res);
    }
</script>

<h1>Donation Page</h1>

<button on:click={donate}>Donate</button>
