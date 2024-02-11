<script lang="ts">
    import {
        Button, CloseButton, Drawer,
        Fileupload, GradientButton, Helper, Input, Label, Range,
    } from "flowbite-svelte";
    import {Alert} from "$lib/pb/alerts_pb";
    import {createAlert} from "$lib/alerts";
    import {uploadToStorage} from "$lib/firebase";
    import {BellActiveSolid} from "flowbite-svelte-icons";
    import {sineIn} from "svelte/easing";
    import {fade} from "svelte/transition";
    import {v4 as uuidv4} from "uuid";

    const addNew = async () => {
        if(!gifFileList || !soundFileList || gifFileList.length === 0 || soundFileList.length === 0) {
            return;
        }

        const gifFile = gifFileList[0];
        const soundFile = soundFileList[0];

        try {
            const id = uuidv4().replace(/-/g, "");
            const gifUrl = await uploadToStorage("gifs", id, gifFile, false); // TODO: change name etc
            const soundUrl = await uploadToStorage("sounds", id, soundFile, false);

            const alert = Alert.fromJson({
                id: id,
                from: startValue,
                to: endValue,
                soundUrl: soundUrl,
                gifUrl: gifUrl,
                template: template
            })
            await createAlert(alert)

            hidden = true;
        } catch (e) {
            console.error(e); // TODO: Show error to user
        }
    }

    let gifFileList: FileList;
    $: gifImage = gifFileList ? URL.createObjectURL(gifFileList[0]) : null;

    let soundFileList: FileList;

    let startValue = 0;
    let endValue = 10;

    let template = "[USER] donated [AMOUNT] [CURRENCY]!";

    let labelClass = "space-y-2";

    let transitionParams = {
        x: 300,
        duration: 200,
        easing: sineIn,
    };

    export let hidden = true;

    let divClass = 'overflow-y-auto z-50 p-4 rounded-xl m-0 md:m-10 bg-gray-100 dark:bg-gray-800';
</script>

<Drawer transitionType="fly" {transitionParams} bind:hidden {divClass} placement="right" width="w-96">
    <div class="flex items-center">
        <h5 id="drawer-label" class="inline-flex items-center mb-6 text-base font-semibold text-gray-500 uppercase dark:text-gray-400">
            <BellActiveSolid class="w-4 h-4 me-2.5" />
            New Alert
        </h5>
        <CloseButton on:click={() => (hidden = true)} class="mb-4 dark:text-white" />
    </div>

    <div class="w-full space-y-4">
        <Label class={labelClass}>From</Label>
        <Input type="number" bind:value={startValue} />

        <Label class={labelClass}>To</Label>
        <Input type="number" bind:value={endValue} />

        <Label class={labelClass}>GIF</Label>
        <Fileupload accept="image/gif" size="sm" bind:files={gifFileList} />
        {#if gifFileList}
            <img src={gifImage} alt="GIF" class="px-8 py-2" in:fade />
        {/if}

        <Label class={labelClass}>Sound</Label>
        <Fileupload size="sm" bind:files={soundFileList} accept="audio/*" />

        <Label class={labelClass}>Template</Label>
        <Input type="text" bind:value={template} />

        <Button on:click={addNew} class="{labelClass} mt-7 w-full">Save</Button>
    </div>

</Drawer>
