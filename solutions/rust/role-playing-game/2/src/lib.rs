// This stub file contains items that aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
#![allow(unused)]

pub struct Player {
    pub health: u32,
    pub mana: Option<u32>,
    pub level: u32,
}

impl Player {
    pub fn revive(&self) -> Option<Player> {
        if self.health == 0 {
            let restored_mana = match self.level {
                10.. => Some(100),
                0.. => None,
            };
            Some(Player {
                health: 100,
                mana: restored_mana,
                level: self.level,
            })
        } else {
            None
        }
    }

    pub fn cast_spell(&mut self, mana_cost: u32) -> u32 {
        match self.mana {
            Some(ref mut mana) => {
                if *mana >= mana_cost {
                    *mana -= mana_cost;
                    mana_cost * 2
                } else {
                    0
                }
            }
            None => {
                self.health = self.health.saturating_sub(mana_cost);
                0
            }
        }
    }
}
