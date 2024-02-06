<script lang="ts">
    import axios from 'axios';
    import {PUBLIC_GET_LISTENERS_LINK, PUBLIC_SENT_DONATE_LINK} from "$env/static/public";
    import {DonateStatus, SendDonateRequest, SendDonateResponse} from "$lib/pb/functions_pb";
    import {
        Card,
    } from "flowbite-svelte";
    import {page} from "$app/stores";
    import {onMount} from "svelte";
    import {FetchError, getAvatarUrl, getPageDocByUsername} from "$lib/page";
    import {goto} from "$app/navigation";
    import Header from "./Header.svelte";
    import {emailStore, usernameStore} from "$lib/stores";
    import Form from "./Form.svelte";

    const sendDonate = async () => {
        if (isBlocked) return;
        if (($usernameStore === "" || $emailStore === "") && !isSignedIn) return;

        const amountFloat = parseFloat(amount);
        if (amountFloat <= 1) return;

        const sdReq = new SendDonateRequest({
            amount: amountFloat,
            currency: currency,
            email: $emailStore,
            sender: $usernameStore,
            recipientId: uid,
            message: message,
            status: DonateStatus.INITIATED
        });

        const res = await axios.post(PUBLIC_SENT_DONATE_LINK, sdReq.toBinary(), { responseType: 'arraybuffer' })
        if(res.status !== 200) {
            console.error(res.data); // TODO: handle error
            return;
        }

        const sdRes = SendDonateResponse.fromBinary(new Uint8Array(res.data));
        window.location.href = sdRes.redirectUrl;
    }

    let recipient = $page.params.recipient;
    let uid:string = "";
    let avatarUrl:string = "";
    let description:string = "";
    let url:string = "";

    let amount:string = "5";
    let message:string = "";
    let currency:string = "PLN";

    let isSignedIn = false;
    $: isBlocked = !isSignedIn && $emailStore === "" // TODO: better validation

    let hasListeners: boolean = true;

    const getListeners = async () => {
        // TODO: error handling
        const res = await axios.get(PUBLIC_GET_LISTENERS_LINK + "?uid=" + uid)
        if(res.status !== 200) {
            console.error(res.data); // TODO: handle error
            return;
        }
        hasListeners = res.data !== 0;
    }

    const fetchPageData = async () => {
        try {
            const doc = await getPageDocByUsername(recipient);
            const data = doc.data()
            if (!data) {
                await goto("/");
                return;
            }

            uid = data.uid;
            getListeners().then(r => {})
            setInterval(getListeners, 10000)

            console.log("uid", uid)

            avatarUrl = await getAvatarUrl(uid);
            description = data.description;
            url = data.url;
        } catch (e) {
            if(e instanceof FetchError) {
                await goto("/");
                return;
            }
            throw e;
        }
    }

    onMount(async () => {
        await fetchPageData()
    });

    $: cardClass = uid === "" ? "blur" : "";
</script>

<div class="flex justify-center items-center h-screen">
    <Card padding="xl" class="{cardClass} w-full sm:w-[75%] md:w-[50%] xl:w-[35%] max-w-none">

        <div class="px-[5%]">
            <Header {url} {recipient} {avatarUrl} {description}/>

            <Form {recipient} {sendDonate} {isBlocked} {hasListeners} bind:amount bind:message bind:currency/>
        </div>
    </Card>
</div>