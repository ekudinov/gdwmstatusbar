#include <alsa/asoundlib.h>
#include <alsa/control.h>

//string to return 
char str[4];

char *
get_vol(void)
{
    long min, max, volume = 0;
    int res, is_muted = 0;
    snd_mixer_t *handle;
    snd_mixer_selem_id_t *sid;
    const char *card = "default";
    const char *selem_name = "Master";
	
    snd_mixer_open(&handle, 0);
    snd_mixer_attach(handle, card);
    snd_mixer_selem_register(handle, NULL, NULL);
    snd_mixer_load(handle);

    snd_mixer_selem_id_alloca(&sid);
    snd_mixer_selem_id_set_index(sid, 0);
    snd_mixer_selem_id_set_name(sid, selem_name);
    snd_mixer_elem_t* elem = snd_mixer_find_selem(handle, sid);

    snd_mixer_selem_get_playback_volume_range(elem, &min, &max);
    snd_mixer_selem_get_playback_volume(elem, 0, &volume);
    snd_mixer_selem_get_playback_switch(elem, 0, &is_muted);
    snd_mixer_close(handle);

    // return "M" when sound is muted
    if (is_muted==0) {
        return "M";
    }
    res = ((double)volume / max) * 100;

    // convert int to str + "%"
    sprintf(str, "%d%%", res);
    return str;
}

